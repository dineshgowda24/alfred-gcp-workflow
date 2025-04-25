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
