module github.com/talos-systems/talos

go 1.16

replace (
	github.com/talos-systems/talos/pkg/machinery => ./pkg/machinery

	// forked go-yaml that introduces RawYAML interface, which can be used to populate YAML fields using bytes
	// which are then encoded as a valid YAML blocks with proper indentiation
	gopkg.in/yaml.v3 => github.com/unix4ever/yaml v0.0.0-20210315173758-8fb30b8e5a5b

	// See https://github.com/talos-systems/go-loadbalancer/pull/4
	// `go get github.com/smira/tcpproxy@combined-fixes`, then copy pseudo-version there
	inet.af/tcpproxy => github.com/smira/tcpproxy v0.0.0-20201015133617-de5f7797b95b
)

require (
	github.com/AlekSi/pointer v1.1.0
	github.com/BurntSushi/toml v0.3.1
	github.com/aws/aws-sdk-go v1.27.0 // indirect
	github.com/beevik/ntp v0.3.0
	github.com/containerd/cgroups v1.0.1
	github.com/containerd/containerd v1.5.2
	github.com/containerd/cri v1.19.0
	github.com/containerd/go-cni v1.0.2
	github.com/containerd/typeurl v1.0.2
	github.com/containernetworking/cni v0.8.1
	github.com/containernetworking/plugins v0.9.1
	github.com/coreos/go-iptables v0.6.0
	github.com/coreos/go-semver v0.3.0
	github.com/cosi-project/runtime v0.0.0-20210603165903-ca95c7538d17
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v20.10.7+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/dustin/go-humanize v1.0.0
	github.com/elazarl/goproxy v0.0.0-20210110162100-a92cc753f88e // indirect
	github.com/emicklei/dot v0.16.0
	github.com/emicklei/go-restful v2.15.0+incompatible // indirect
	github.com/evanphx/json-patch v4.9.0+incompatible
	github.com/fatih/color v1.12.0
	github.com/firecracker-microvm/firecracker-go-sdk v0.22.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fullsailor/pkcs7 v0.0.0-20190404230743-d7302db945fa
	github.com/gdamore/tcell/v2 v2.3.11
	github.com/gizak/termui/v3 v3.1.0
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.2.0
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/go-getter v1.5.3
	github.com/hashicorp/go-multierror v1.1.1
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/insomniacslk/dhcp v0.0.0-20210528123148-fb4eaaa00ad2
	github.com/jsimonetti/rtnetlink v0.0.0-20210531051304-b34cb89a106b
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-isatty v0.0.13
	github.com/mdlayher/arp v0.0.0-20191213142603-f72070a231fc
	github.com/mdlayher/ethtool v0.0.0-20210210192532-2b88debcdd43
	github.com/mdlayher/genetlink v1.0.0
	github.com/mdlayher/netlink v1.4.1
	github.com/mdlayher/raw v0.0.0-20210412142147-51b895745faf // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20200929063507-e6143ca7d51d
	github.com/pelletier/go-toml v1.9.0 // indirect
	github.com/pin/tftp v2.1.0+incompatible
	github.com/prometheus/procfs v0.6.0
	github.com/rivo/tview v0.0.0-20210531104647-807e706f86d1
	github.com/rs/xid v1.3.0
	github.com/ryanuber/columnize v2.1.2+incompatible
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/smira/go-xz v0.0.0-20201019130106-9921ed7a9935
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/talos-systems/crypto v0.2.1-0.20210601174604-cd18ef62eb9f
	github.com/talos-systems/go-blockdevice v0.2.1-0.20210526155905-30c2bc3cb62a
	github.com/talos-systems/go-cmd v0.0.0-20210216164758-68eb0067e0f0
	github.com/talos-systems/go-debug v0.2.1-0.20210525175311-3d0a6e1bf5e3
	github.com/talos-systems/go-kmsg v0.1.0
	github.com/talos-systems/go-loadbalancer v0.1.1
	github.com/talos-systems/go-procfs v0.0.0-20210108152626-8cbc42d3dc24
	github.com/talos-systems/go-retry v0.3.1-0.20210616175710-c78cc953d9e9
	github.com/talos-systems/go-smbios v0.0.0-20210422124317-d3a32bea731a
	github.com/talos-systems/grpc-proxy v0.2.0
	github.com/talos-systems/net v0.2.1-0.20210212213224-05190541b0fa
	github.com/talos-systems/talos/pkg/machinery v0.0.0-00010101000000-000000000000
	github.com/u-root/u-root v7.0.0+incompatible
	github.com/vishvananda/netns v0.0.0-20210104183010-2eb08e3e575f // indirect
	github.com/vmware-tanzu/sonobuoy v0.51.0
	github.com/vmware/govmomi v0.26.0
	github.com/vmware/vmw-guestinfo v0.0.0-20200218095840-687661b8bd8e
	go.etcd.io/etcd/api/v3 v3.5.0-rc.1
	go.etcd.io/etcd/client/pkg/v3 v3.5.0-rc.1
	go.etcd.io/etcd/client/v3 v3.5.0-rc.1
	go.etcd.io/etcd/etcdutl/v3 v3.5.0-rc.1
	go.uber.org/zap v1.17.0
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	golang.org/x/oauth2 v0.0.0-20210427180440-81ed05c6b58c // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20210506160403-92e472f520a5
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/freddierice/go-losetup.v1 v1.0.0-20170407175016-fc9adea44124
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	inet.af/netaddr v0.0.0-20210430201628-1d252cf8125e
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/apiserver v0.21.1 // indirect
	k8s.io/client-go v0.21.1
	k8s.io/cri-api v0.21.1
	k8s.io/kubectl v0.21.1
	k8s.io/kubelet v0.21.1
	k8s.io/utils v0.0.0-20210305010621-2afb4311ab10 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.1 // indirect
)
