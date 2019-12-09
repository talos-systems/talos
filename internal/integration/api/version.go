// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// +build integration_api

package api

import (
	"context"
	"time"

	"github.com/talos-systems/talos/api/machine"
	"github.com/talos-systems/talos/cmd/osctl/pkg/client"
	"github.com/talos-systems/talos/internal/integration/base"
	"github.com/talos-systems/talos/pkg/retry"
)

// VersionSuite verifies version API
type VersionSuite struct {
	base.APISuite

	ctx       context.Context
	ctxCancel context.CancelFunc
}

// SuiteName ...
func (suite *VersionSuite) SuiteName() string {
	return "api.VersionSuite"
}

// SetupTest ...
func (suite *VersionSuite) SetupTest() {
	// make sure API calls have timeout
	suite.ctx, suite.ctxCancel = context.WithTimeout(context.Background(), 2*time.Minute)
}

// TearDownTest ...
func (suite *VersionSuite) TearDownTest() {
	suite.ctxCancel()
}

// TestExpectedVersionMaster verifies master node version matches expected
func (suite *VersionSuite) TestExpectedVersionMaster() {
	v, err := suite.Client.Version(suite.ctx)
	suite.Require().NoError(err)

	suite.Assert().Equal(suite.Version, v.Response[0].Version.Tag)
}

// TestSameVersionCluster verifies that all the nodes are on the same version
func (suite *VersionSuite) TestSameVersionCluster() {
	nodes := suite.DiscoverNodes()
	suite.Require().NotEmpty(nodes)

	ctx := client.WithNodes(suite.ctx, nodes...)

	var v *machine.VersionReply

	err := retry.Constant(
		time.Minute,
	).Retry(func() error {
		var e error
		v, e = suite.Client.Version(ctx)

		return retry.ExpectedError(e)
	})

	suite.Require().NoError(err)

	suite.Require().Len(v.Response, len(nodes))

	expectedVersion := v.Response[0].Version.Tag
	for _, version := range v.Response {
		suite.Assert().Equal(expectedVersion, version.Version.Tag)
	}
}

func init() {
	allSuites = append(allSuites, new(VersionSuite))
}
