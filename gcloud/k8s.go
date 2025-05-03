package gcloud

type K8sCluster struct {
	Name                 string `json:"name"`
	Status               string `json:"status"`
	Location             string `json:"location"`
	CurrentMasterVersion string `json:"currentMasterVersion"`
	CurrentNodeCount     int    `json:"currentNodeCount"`
	CreatedAt            string `json:"createTime"`
}

func (c K8sCluster) IsRegionSupported(config *Config) bool {
	return isComputeRegionSupported(config)
}

func ListK8sClusters(config *Config) ([]K8sCluster, error) {
	args := []string{"clusters", "list", "--format=json(name,status,location,currentMasterVersion,currentNodeCount,createTime)"}
	if config != nil && config.GetRegionName() != "" {
		args = append(args, "--region="+config.GetRegionName())
	}
	return runGCloudCmd[[]K8sCluster](config, "container", args...)
}
