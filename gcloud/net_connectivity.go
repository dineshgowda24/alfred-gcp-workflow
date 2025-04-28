package gcloud

// https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels/list
type VPNTunnel struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Region     string `json:"region"`
	VPNGateway string `json:"vpnGateway"`
}

func ListVPNTunnels(config *Config) ([]VPNTunnel, error) {
	return runGCloudCmd[[]VPNTunnel](
		config,
		"compute", "vpn-tunnels", "list",
		"--format=json(name,status,region,vpnGateway)",
	)
}
