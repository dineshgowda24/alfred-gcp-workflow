package gcloud

type VPCNetwork struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	CreationTime         string `json:"creationTimestamp"`
	XCloudBgpRoutingMode string `json:"x_gcloud_bgp_routing_mode"`
}

func ListVPCNetworks(config *Config) ([]VPCNetwork, error) {
	return runGCloudCmd[[]VPCNetwork](
		config,
		"compute", "networks", "list",
		"--format=json(id,name,description,creationTimestamp,x_gcloud_bgp_routing_mode)",
	)
}

type VPCRoute struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Network      string `json:"network"`
	Priority     int    `json:"priority"`
	CreationTime string `json:"creationTimestamp"`
}

func ListVPCRoutes(config *Config) ([]VPCRoute, error) {
	return runGCloudCmd[[]VPCRoute](
		config,
		"compute", "routes", "list",
		"--format=json(id,name,network,priority,creationTimestamp)",
	)
}
