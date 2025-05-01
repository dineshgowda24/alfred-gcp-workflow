package gcloud

// https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels
type VPNTunnel struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Region     string `json:"region"`
	VPNGateway string `json:"vpnGateway"`
}

// https://cloud.google.com/compute/docs/reference/rest/v1/vpnGateways
type VPNGateway struct {
	Name              string `json:"name"`
	GatewayIPVersion  string `json:"gatewayIpVersion"`
	Network           string `json:"network"`
	Region            string `json:"region"`
	CreationTimestamp string `json:"creationTimestamp"`
}

// https://cloud.google.com/compute/docs/reference/rest/v1/routers
type CloudRouter struct {
	Name              string `json:"name"`
	Network           string `json:"network"`
	Region            string `json:"region"`
	CreationTimestamp string `json:"creationTimestamp"`
}

func ListVPNTunnels(config *Config) ([]VPNTunnel, error) {
	return runGCloudCmd[[]VPNTunnel](
		config,
		"compute", "vpn-tunnels", "list",
		"--format=json(name,status,region,vpnGateway)",
	)
}

func ListVPNGateways(config *Config) ([]VPNGateway, error) {
	return runGCloudCmd[[]VPNGateway](
		config,
		"compute", "vpn-gateways", "list",
		"--format=json(name,gatewayIpVersion,network,region,creationTimestamp)",
	)
}

func ListCloudRouters(config *Config) ([]CloudRouter, error) {
	return runGCloudCmd[[]CloudRouter](
		config,
		"compute", "routers", "list",
		"--format=json(name,network,region,creationTimestamp)",
	)
}
