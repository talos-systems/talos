// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// +build integration_provision

package provision

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/stretchr/testify/suite"

	machineapi "github.com/talos-systems/talos/api/machine"
	"github.com/talos-systems/talos/cmd/talosctl/pkg/mgmt/helpers"
	"github.com/talos-systems/talos/internal/app/machined/pkg/runtime"
	"github.com/talos-systems/talos/internal/integration/base"
	"github.com/talos-systems/talos/internal/pkg/cluster/check"
	"github.com/talos-systems/talos/internal/pkg/provision"
	"github.com/talos-systems/talos/internal/pkg/provision/access"
	"github.com/talos-systems/talos/internal/pkg/provision/providers/firecracker"
	talosclient "github.com/talos-systems/talos/pkg/client"
	"github.com/talos-systems/talos/pkg/config"
	"github.com/talos-systems/talos/pkg/config/types/v1alpha1"
	"github.com/talos-systems/talos/pkg/config/types/v1alpha1/generate"
	"github.com/talos-systems/talos/pkg/constants"
	talosnet "github.com/talos-systems/talos/pkg/net"
	"github.com/talos-systems/talos/pkg/retry"
)

type upgradeSpec struct {
	ShortName string

	SourceKernelPath     string
	SourceInitramfsPath  string
	SourceInstallerImage string
	SourceVersion        string

	TargetInstallerImage string
	TargetVersion        string

	MasterNodes int
	WorkerNodes int

	UpgradePreserve bool
}

const (
	stableVersion = "v0.5.1"
	nextVersion   = "v0.6.0-alpha.6"
)

var (
	defaultNameservers = []net.IP{net.ParseIP("8.8.8.8"), net.ParseIP("1.1.1.1")}
	defaultCNIBinPath  = []string{"/opt/cni/bin"}
)

const (
	defaultCNIConfDir  = "/etc/cni/conf.d"
	defaultCNICacheDir = "/var/lib/cni"
)

func trimVersion(version string) string {
	// remove anything extra after semantic version core, `v0.3.2-1-abcd` -> `v0.3.2`
	return regexp.MustCompile(`(-\d+-g[0-9a-f]+)$`).ReplaceAllString(version, "")
}

// upgradeBetweenTwoLastReleases upgrades between two last releases of Talos.
func upgradeBetweenTwoLastReleases() upgradeSpec {
	return upgradeSpec{
		ShortName: fmt.Sprintf("%s-%s", stableVersion, nextVersion),

		SourceKernelPath:     helpers.ArtifactPath(filepath.Join(trimVersion(stableVersion), constants.KernelAsset)),
		SourceInitramfsPath:  helpers.ArtifactPath(filepath.Join(trimVersion(stableVersion), constants.InitramfsAsset)),
		SourceInstallerImage: fmt.Sprintf("%s:%s", constants.DefaultInstallerImageRepository, stableVersion),
		SourceVersion:        stableVersion,

		TargetInstallerImage: fmt.Sprintf("%s:%s", constants.DefaultInstallerImageRepository, nextVersion),
		TargetVersion:        nextVersion,

		MasterNodes: DefaultSettings.MasterNodes,
		WorkerNodes: DefaultSettings.WorkerNodes,
	}
}

// upgradeLastReleaseToCurrent upgrades last release to the current version of Talos.
func upgradeLastReleaseToCurrent() upgradeSpec {
	return upgradeSpec{
		ShortName: fmt.Sprintf("%s-%s", nextVersion, DefaultSettings.CurrentVersion),

		SourceKernelPath:     helpers.ArtifactPath(filepath.Join(trimVersion(nextVersion), constants.KernelAsset)),
		SourceInitramfsPath:  helpers.ArtifactPath(filepath.Join(trimVersion(nextVersion), constants.InitramfsAsset)),
		SourceInstallerImage: fmt.Sprintf("%s:%s", constants.DefaultInstallerImageRepository, nextVersion),
		SourceVersion:        nextVersion,

		TargetInstallerImage: fmt.Sprintf("%s/%s:%s", DefaultSettings.TargetInstallImageRegistry, constants.DefaultInstallerImageName, DefaultSettings.CurrentVersion),
		TargetVersion:        DefaultSettings.CurrentVersion,

		MasterNodes: DefaultSettings.MasterNodes,
		WorkerNodes: DefaultSettings.WorkerNodes,
	}
}

// upgradeSingeNodePreserve upgrade last release of Talos to the current version of Talos for single-node cluster with preserve.
func upgradeSingeNodePreserve() upgradeSpec {
	return upgradeSpec{
		ShortName: fmt.Sprintf("preserve-%s-%s", nextVersion, DefaultSettings.CurrentVersion),

		SourceKernelPath:     helpers.ArtifactPath(filepath.Join(trimVersion(nextVersion), constants.KernelAsset)),
		SourceInitramfsPath:  helpers.ArtifactPath(filepath.Join(trimVersion(nextVersion), constants.InitramfsAsset)),
		SourceInstallerImage: fmt.Sprintf("%s:%s", constants.DefaultInstallerImageRepository, nextVersion),
		SourceVersion:        nextVersion,

		TargetInstallerImage: fmt.Sprintf("%s/%s:%s", DefaultSettings.TargetInstallImageRegistry, constants.DefaultInstallerImageName, DefaultSettings.CurrentVersion),
		TargetVersion:        DefaultSettings.CurrentVersion,

		MasterNodes:     1,
		WorkerNodes:     0,
		UpgradePreserve: true,
	}
}

type UpgradeSuite struct {
	suite.Suite
	base.TalosSuite

	specGen func() upgradeSpec
	spec    upgradeSpec

	track int

	provisioner provision.Provisioner

	configBundle *v1alpha1.ConfigBundle

	clusterAccess *access.Adapter

	ctx       context.Context
	ctxCancel context.CancelFunc

	stateDir string
}

// SetupSuite ...
func (suite *UpgradeSuite) SetupSuite() {
	// call generate late in the flow, as it needs to pick up settings overridden by test runner
	suite.spec = suite.specGen()

	suite.T().Logf("upgrade spec = %v", suite.spec)

	// timeout for the whole test
	suite.ctx, suite.ctxCancel = context.WithTimeout(context.Background(), 30*time.Minute)

	var err error

	suite.provisioner, err = firecracker.NewProvisioner(suite.ctx)
	suite.Require().NoError(err)
}

// TearDownSuite ...
func (suite *UpgradeSuite) TearDownSuite() {
	if suite.T().Failed() && suite.Cluster != nil {
		// for failed tests, produce crash dump for easier debugging,
		// as cluster is going to be torn down below
		suite.provisioner.CrashDump(suite.ctx, suite.Cluster, os.Stderr)

		if suite.clusterAccess != nil {
			suite.clusterAccess.CrashDump(suite.ctx, os.Stderr)
		}
	}

	if suite.clusterAccess != nil {
		suite.Assert().NoError(suite.clusterAccess.Close())
	}

	if suite.Cluster != nil {
		suite.Assert().NoError(suite.provisioner.Destroy(suite.ctx, suite.Cluster))
	}

	suite.ctxCancel()

	if suite.stateDir != "" {
		suite.Assert().NoError(os.RemoveAll(suite.stateDir))
	}

	if suite.provisioner != nil {
		suite.Assert().NoError(suite.provisioner.Close())
	}
}

// setupCluster provisions source clusters and waits for health.
func (suite *UpgradeSuite) setupCluster() {
	shortNameHash := sha256.Sum256([]byte(suite.spec.ShortName))
	clusterName := fmt.Sprintf("upgrade.%x", shortNameHash[:8])

	_, cidr, err := net.ParseCIDR(DefaultSettings.CIDR)
	suite.Require().NoError(err)

	var gatewayIP net.IP

	gatewayIP, err = talosnet.NthIPInNetwork(cidr, 1)
	suite.Require().NoError(err)

	ips := make([]net.IP, suite.spec.MasterNodes+suite.spec.WorkerNodes)

	for i := range ips {
		ips[i], err = talosnet.NthIPInNetwork(cidr, i+2)
		suite.Require().NoError(err)
	}

	suite.stateDir, err = ioutil.TempDir("", "talos-integration")
	suite.Require().NoError(err)

	suite.T().Logf("initalizing provisioner with cluster name %q, state directory %q", clusterName, suite.stateDir)

	request := provision.ClusterRequest{
		Name: clusterName,

		Network: provision.NetworkRequest{
			Name:        clusterName,
			CIDR:        *cidr,
			GatewayAddr: gatewayIP,
			MTU:         DefaultSettings.MTU,
			Nameservers: defaultNameservers,
			CNI: provision.CNIConfig{
				BinPath:  defaultCNIBinPath,
				ConfDir:  defaultCNIConfDir,
				CacheDir: defaultCNICacheDir,
			},
		},

		KernelPath:    suite.spec.SourceKernelPath,
		InitramfsPath: suite.spec.SourceInitramfsPath,

		SelfExecutable: suite.TalosctlPath,
		StateDirectory: suite.stateDir,
	}

	defaultInternalLB, _ := suite.provisioner.GetLoadBalancers(request.Network)

	genOptions := suite.provisioner.GenOptions(request.Network)

	for _, registryMirror := range DefaultSettings.RegistryMirrors {
		parts := strings.SplitN(registryMirror, "=", 2)
		suite.Require().Len(parts, 2)

		genOptions = append(genOptions, generate.WithRegistryMirror(parts[0], parts[1]))
	}

	masterEndpoints := make([]string, suite.spec.MasterNodes)
	for i := range masterEndpoints {
		masterEndpoints[i] = ips[i].String()
	}

	suite.configBundle, err = config.NewConfigBundle(config.WithInputOptions(
		&config.InputOptions{
			ClusterName: clusterName,
			Endpoint:    fmt.Sprintf("https://%s:6443", defaultInternalLB),
			KubeVersion: "", // keep empty so that default version is used per Talos version
			GenOptions: append(
				genOptions,
				generate.WithEndpointList(masterEndpoints),
				generate.WithInstallImage(suite.spec.SourceInstallerImage),
			),
		}))
	suite.Require().NoError(err)

	for i := 0; i < suite.spec.MasterNodes; i++ {
		var cfg runtime.Configurator

		if i == 0 {
			cfg = suite.configBundle.Init()
		} else {
			cfg = suite.configBundle.ControlPlane()
		}

		request.Nodes = append(request.Nodes,
			provision.NodeRequest{
				Name:     fmt.Sprintf("master-%d", i+1),
				IP:       ips[i],
				Memory:   DefaultSettings.MemMB * 1024 * 1024,
				NanoCPUs: DefaultSettings.CPUs * 1000 * 1000 * 1000,
				DiskSize: DefaultSettings.DiskGB * 1024 * 1024 * 1024,
				Config:   cfg,
			})
	}

	for i := 1; i <= suite.spec.WorkerNodes; i++ {
		request.Nodes = append(request.Nodes,
			provision.NodeRequest{
				Name:     fmt.Sprintf("worker-%d", i),
				IP:       ips[suite.spec.MasterNodes+i-1],
				Memory:   DefaultSettings.MemMB * 1024 * 1024,
				NanoCPUs: DefaultSettings.CPUs * 1000 * 1000 * 1000,
				DiskSize: DefaultSettings.DiskGB * 1024 * 1024 * 1024,
				Config:   suite.configBundle.Join(),
			})
	}

	suite.Cluster, err = suite.provisioner.Create(suite.ctx, request, provision.WithBootlader(true), provision.WithTalosConfig(suite.configBundle.TalosConfig()))
	suite.Require().NoError(err)

	suite.clusterAccess = access.NewAdapter(suite.Cluster, provision.WithTalosConfig(suite.configBundle.TalosConfig()))

	suite.waitForClusterHealth()
}

// waitForClusterHealth asserts cluster health after any change.
func (suite *UpgradeSuite) waitForClusterHealth() {
	checkCtx, checkCtxCancel := context.WithTimeout(suite.ctx, 10*time.Minute)
	defer checkCtxCancel()

	suite.Require().NoError(check.Wait(checkCtx, suite.clusterAccess, check.DefaultClusterChecks(), check.StderrReporter()))
}

func (suite *UpgradeSuite) assertSameVersionCluster(client *talosclient.Client, expectedVersion string) {
	nodes := make([]string, len(suite.Cluster.Info().Nodes))

	for i, node := range suite.Cluster.Info().Nodes {
		nodes[i] = node.PrivateIP.String()
	}

	ctx := talosclient.WithNodes(suite.ctx, nodes...)

	var v *machineapi.VersionResponse

	err := retry.Constant(
		time.Minute,
	).Retry(func() error {
		var e error
		v, e = client.Version(ctx)

		return retry.ExpectedError(e)
	})

	suite.Require().NoError(err)

	suite.Require().Len(v.Messages, len(nodes))

	for _, version := range v.Messages {
		suite.Assert().Equal(expectedVersion, version.Version.Tag)
	}
}

func (suite *UpgradeSuite) readVersion(client *talosclient.Client, nodeCtx context.Context) (version string, err error) {
	var v *machineapi.VersionResponse

	v, err = client.Version(nodeCtx)
	if err != nil {
		return
	}

	version = v.Messages[0].Version.Tag

	return
}

func (suite *UpgradeSuite) upgradeNode(client *talosclient.Client, node provision.NodeInfo) {
	suite.T().Logf("upgrading node %s", node.PrivateIP)

	nodeCtx := talosclient.WithNodes(suite.ctx, node.PrivateIP.String())

	resp, err := client.Upgrade(nodeCtx, suite.spec.TargetInstallerImage, suite.spec.UpgradePreserve)
	suite.Require().NoError(err)

	suite.Require().Equal("Upgrade request received", resp.Messages[0].Ack)

	// wait for the version to be equal to target version
	suite.Require().NoError(retry.Constant(5 * time.Minute).Retry(func() error {
		var version string

		version, err = suite.readVersion(client, nodeCtx)
		if err != nil {
			// API might be unresponsive during upgrade
			return retry.ExpectedError(err)
		}

		if version != suite.spec.TargetVersion {
			// upgrade not finished yet
			return retry.ExpectedError(fmt.Errorf("node %q version doesn't match expected: expected %q, got %q", node.PrivateIP.String(), suite.spec.TargetVersion, version))
		}

		return nil
	}))

	suite.waitForClusterHealth()
}

// TestRolling performs rolling upgrade starting with master nodes.
func (suite *UpgradeSuite) TestRolling() {
	suite.setupCluster()

	client, err := suite.clusterAccess.Client()
	suite.Require().NoError(err)

	// verify initial cluster version
	suite.assertSameVersionCluster(client, suite.spec.SourceVersion)

	// upgrade master nodes
	for _, node := range suite.Cluster.Info().Nodes {
		if node.Type == runtime.MachineTypeInit || node.Type == runtime.MachineTypeControlPlane {
			suite.upgradeNode(client, node)
		}
	}

	// upgrade worker nodes
	for _, node := range suite.Cluster.Info().Nodes {
		if node.Type == runtime.MachineTypeJoin {
			suite.upgradeNode(client, node)
		}
	}

	// verify final cluster version
	suite.assertSameVersionCluster(client, suite.spec.TargetVersion)
}

// SuiteName ...
func (suite *UpgradeSuite) SuiteName() string {
	if suite.spec.ShortName == "" {
		suite.spec = suite.specGen()
	}

	return fmt.Sprintf("provision.UpgradeSuite.%s-TR%d", suite.spec.ShortName, suite.track)
}

func init() {
	allSuites = append(allSuites,
		&UpgradeSuite{specGen: upgradeBetweenTwoLastReleases, track: 0},
		&UpgradeSuite{specGen: upgradeLastReleaseToCurrent, track: 1},
		&UpgradeSuite{specGen: upgradeSingeNodePreserve, track: 0},
	)
}
