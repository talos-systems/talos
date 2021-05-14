// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//nolint:dupl
package network_test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/cosi-project/runtime/pkg/controller/runtime"
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/state"
	"github.com/cosi-project/runtime/pkg/state/impl/inmem"
	"github.com/cosi-project/runtime/pkg/state/impl/namespaced"
	"github.com/jsimonetti/rtnetlink"
	"github.com/stretchr/testify/suite"
	"github.com/talos-systems/go-retry/retry"
	"inet.af/netaddr"

	netctrl "github.com/talos-systems/talos/internal/app/machined/pkg/controllers/network"
	"github.com/talos-systems/talos/pkg/logging"
	"github.com/talos-systems/talos/pkg/resources/network"
	"github.com/talos-systems/talos/pkg/resources/network/nethelpers"
)

type RouteSpecSuite struct {
	suite.Suite

	state state.State

	runtime *runtime.Runtime
	wg      sync.WaitGroup

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (suite *RouteSpecSuite) SetupTest() {
	suite.ctx, suite.ctxCancel = context.WithTimeout(context.Background(), 3*time.Minute)

	suite.state = state.WrapCore(namespaced.NewState(inmem.Build))

	var err error

	suite.runtime, err = runtime.NewRuntime(suite.state, logging.Wrap(log.Writer()))
	suite.Require().NoError(err)

	suite.Require().NoError(suite.runtime.RegisterController(&netctrl.RouteSpecController{}))

	suite.startRuntime()
}

func (suite *RouteSpecSuite) startRuntime() {
	suite.wg.Add(1)

	go func() {
		defer suite.wg.Done()

		suite.Assert().NoError(suite.runtime.Run(suite.ctx))
	}()
}

func (suite *RouteSpecSuite) assertRoute(destination netaddr.IPPrefix, gateway netaddr.IP, check func(rtnetlink.RouteMessage) error) error {
	conn, err := rtnetlink.Dial(nil)
	suite.Require().NoError(err)

	defer conn.Close() //nolint:errcheck

	routes, err := conn.Route.List()
	suite.Require().NoError(err)

	matching := 0

	for _, route := range routes {
		if !gateway.IPAddr().IP.Equal(route.Attributes.Gateway) {
			continue
		}

		if route.DstLength != destination.Bits {
			continue
		}

		if !destination.IP.IPAddr().IP.Equal(route.Attributes.Dst) {
			continue
		}

		matching++

		if err = check(route); err != nil {
			return retry.ExpectedError(err)
		}
	}

	switch {
	case matching == 1:
		return nil
	case matching == 0:
		return retry.ExpectedError(fmt.Errorf("route to %s via %s not found", destination, gateway))
	default:
		return retry.ExpectedError(fmt.Errorf("route to %s via %s found %d matches", destination, gateway, matching))
	}
}

func (suite *RouteSpecSuite) assertNoRoute(destination netaddr.IPPrefix, gateway netaddr.IP) error {
	conn, err := rtnetlink.Dial(nil)
	suite.Require().NoError(err)

	defer conn.Close() //nolint:errcheck

	routes, err := conn.Route.List()
	suite.Require().NoError(err)

	for _, route := range routes {
		if gateway.IPAddr().IP.Equal(route.Attributes.Gateway) && destination.Bits == route.DstLength && destination.IP.IPAddr().IP.Equal(route.Attributes.Dst) {
			return retry.ExpectedError(fmt.Errorf("route to %s via %s is present", destination, gateway))
		}
	}

	return nil
}

func (suite *RouteSpecSuite) TestLoopback() {
	loopback := network.NewRouteSpec(network.NamespaceName, "loopback")
	*loopback.Status() = network.RouteSpecSpec{
		Family:      nethelpers.FamilyInet4,
		Destination: netaddr.MustParseIPPrefix("127.0.11.0/24"),
		Gateway:     netaddr.MustParseIP("127.0.11.1"),
		OutLinkName: "lo",
		Scope:       nethelpers.ScopeGlobal,
		Table:       nethelpers.TableMain,
		Protocol:    nethelpers.ProtocolStatic,
		Type:        nethelpers.TypeUnicast,
		Layer:       network.ConfigMachineConfiguration,
	}

	for _, res := range []resource.Resource{loopback} {
		suite.Require().NoError(suite.state.Create(suite.ctx, res), "%v", res.Spec())
	}

	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		func() error {
			return suite.assertRoute(netaddr.MustParseIPPrefix("127.0.11.0/24"), netaddr.MustParseIP("127.0.11.1"), func(route rtnetlink.RouteMessage) error {
				suite.Assert().EqualValues(0, route.Attributes.Priority)

				return nil
			})
		}))

	// teardown the route
	for {
		ready, err := suite.state.Teardown(suite.ctx, loopback.Metadata())
		suite.Require().NoError(err)

		if ready {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}

	// torn down address should be removed immediately
	suite.Assert().NoError(suite.assertNoRoute(netaddr.MustParseIPPrefix("127.0.11.0/24"), netaddr.MustParseIP("127.0.11.1")))

	suite.Require().NoError(suite.state.Destroy(suite.ctx, loopback.Metadata()))
}

func (suite *RouteSpecSuite) TestDefaultRoute() {
	// adding default route with high metric to avoid messing up with the actual default route
	def := network.NewRouteSpec(network.NamespaceName, "default")
	*def.Status() = network.RouteSpecSpec{
		Family:      nethelpers.FamilyInet4,
		Destination: netaddr.IPPrefix{},
		Gateway:     netaddr.MustParseIP("127.0.11.2"),
		Scope:       nethelpers.ScopeGlobal,
		Table:       nethelpers.TableMain,
		OutLinkName: "lo",
		Protocol:    nethelpers.ProtocolStatic,
		Type:        nethelpers.TypeUnicast,
		Priority:    1048576,
		Layer:       network.ConfigMachineConfiguration,
	}

	for _, res := range []resource.Resource{def} {
		suite.Require().NoError(suite.state.Create(suite.ctx, res), "%v", res.Spec())
	}

	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		func() error {
			return suite.assertRoute(netaddr.IPPrefix{}, netaddr.MustParseIP("127.0.11.2"), func(route rtnetlink.RouteMessage) error {
				suite.Assert().Nil(route.Attributes.Dst)
				suite.Assert().EqualValues(1048576, route.Attributes.Priority)

				return nil
			})
		}))

	// update the route metric
	_, err := suite.state.UpdateWithConflicts(suite.ctx, def.Metadata(), func(r resource.Resource) error {
		defR := r.(*network.RouteSpec) //nolint:forcetypeassert,errcheck

		defR.Status().Priority = 1048577

		return nil
	})
	suite.Assert().NoError(err)

	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		func() error {
			return suite.assertRoute(netaddr.IPPrefix{}, netaddr.MustParseIP("127.0.11.2"), func(route rtnetlink.RouteMessage) error {
				suite.Assert().Nil(route.Attributes.Dst)

				if route.Attributes.Priority != 1048577 {
					return fmt.Errorf("route metric wasn't updated: %d", route.Attributes.Priority)
				}

				return nil
			})
		}))

	// teardown the route
	for {
		ready, err := suite.state.Teardown(suite.ctx, def.Metadata())
		suite.Require().NoError(err)

		if ready {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}

	// torn down route should be removed immediately
	suite.Assert().NoError(suite.assertNoRoute(netaddr.IPPrefix{}, netaddr.MustParseIP("127.0.11.2")))

	suite.Require().NoError(suite.state.Destroy(suite.ctx, def.Metadata()))
}

func TestRouteSpecSuite(t *testing.T) {
	suite.Run(t, new(RouteSpecSuite))
}
