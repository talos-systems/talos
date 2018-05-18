package etc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"text/template"

	"github.com/autonomy/dianemo/initramfs/pkg/userdata"
	"github.com/autonomy/dianemo/initramfs/pkg/version"
)

const hostsTemplate = `
127.0.0.1       localhost
127.0.0.1       {{ .Hostname }}
{{ .IP }}       {{ .Hostname }}
::1             localhost ip6-localhost ip6-loopback
ff02::1         ip6-allnodes
ff02::2         ip6-allrouters
`

const resolvConfTemplate = `
{{ range $_, $ip := . }}
nameserver {{ $ip }}
{{ end }}
`

const osReleaseTemplate = `
NAME="{{ .Name }}"
ID={{ .ID }}
VERSION_ID={{ .Version }}
PRETTY_NAME="{{ .Name }} ({{ .Version }}) by Autonomy"
HOME_URL="https://dianemo.autonomy.io/"
BUG_REPORT_URL="https://github.com/autonomy/dianemo/issues"
`

// Hosts renders a valid /etc/hosts file and writes it to disk.
func Hosts(s, hostname, ip string) (err error) {
	data := struct {
		IP       string
		Hostname string
	}{
		IP:       ip,
		Hostname: hostname,
	}

	tmpl, err := template.New("").Parse(hostsTemplate)
	if err != nil {
		return
	}
	var buf []byte
	writer := bytes.NewBuffer(buf)
	err = tmpl.Execute(writer, data)
	if err != nil {
		return
	}

	if err := ioutil.WriteFile(path.Join(s, "/etc/hosts"), writer.Bytes(), 0644); err != nil {
		return fmt.Errorf("write /etc/hosts: %v", err)
	}

	return nil
}

// ResolvConf renders a valid /etc/resolv.conf file and writes it to disk.
func ResolvConf(s string, userdata userdata.UserData) (err error) {
	tmpl, err := template.New("").Parse(resolvConfTemplate)
	if err != nil {
		return
	}
	var buf []byte
	writer := bytes.NewBuffer(buf)
	err = tmpl.Execute(writer, userdata.OS.Network.Nameservers)
	if err != nil {
		return
	}

	if err := ioutil.WriteFile(path.Join(s, "/etc/resolv.conf"), writer.Bytes(), 0644); err != nil {
		return fmt.Errorf("write /etc/resolv.conf: %v", err)
	}

	return nil
}

// OSRelease renders a valid /etc/os-release file and writes it to disk. The
// node's OS Image field is reported by the node from /etc/os-release.
func OSRelease(s string) (err error) {
	data := struct {
		Name    string
		ID      string
		Version string
	}{
		Name:    version.Name,
		ID:      strings.ToLower(version.Name),
		Version: version.Tag,
	}

	tmpl, err := template.New("").Parse(osReleaseTemplate)
	if err != nil {
		return
	}
	var buf []byte
	writer := bytes.NewBuffer(buf)
	err = tmpl.Execute(writer, data)
	if err != nil {
		return
	}

	if err := ioutil.WriteFile(path.Join(s, "/etc/os-release"), writer.Bytes(), 0644); err != nil {
		return fmt.Errorf("write /etc/os-release: %v", err)
	}

	return nil
}
