/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package iso

import (
	"net"

	"github.com/talos-systems/talos/internal/pkg/runtime"
	"github.com/talos-systems/talos/pkg/crypto/x509"
	"github.com/talos-systems/talos/pkg/userdata"
)

// ISO is a platform for installing Talos via an ISO image.
type ISO struct{}

// Name implements the platform.Platform interface.
func (i *ISO) Name() string {
	return "ISO"
}

// UserData implements the platform.Platform interface.
func (i *ISO) UserData() (data *userdata.UserData, err error) {
	data = &userdata.UserData{
		Security: &userdata.Security{
			OS: &userdata.OSSecurity{
				CA: &x509.PEMEncodedCertificateAndKey{},
			},
			Kubernetes: &userdata.KubernetesSecurity{
				CA: &x509.PEMEncodedCertificateAndKey{},
			},
		},
		Install: &userdata.Install{
			Force:      true,
			Disk:       "/dev/sda",
			Bootloader: true,
		},
	}

	return data, nil
}

// Mode implements the platform.Platform interface.
func (i *ISO) Mode() runtime.Mode {
	return runtime.Interactive
}

// Hostname implements the platform.Platform interface.
func (i *ISO) Hostname() (hostname []byte, err error) {
	return nil, nil
}

// ExternalIPs provides any external addresses assigned to the instance
func (i *ISO) ExternalIPs() (addrs []net.IP, err error) {
	return addrs, err
}
