package gcloud

type K8sCluster struct {
	Name string `json:"name"`
	// one of [STATUS_UNSPECIFIED PROVISIONING RUNNING RECONCILING STOPPING ERROR DEGRADED]
	Status string `json:"status"`
	// Format 2023-11-01T09:09:29+00:00
	CreatedAt            string `json:"createTime"`
	Location             string `json:"location"`
	CurrentMasterVersion string `json:"currentMasterVersion"`
	CurrentNodeCount     int    `json:"currentNodeCount"`
}

func ListK8sClusters(config *Config) ([]K8sCluster, error) {
	return runGCloudCmd[[]K8sCluster](
		config,
		"container", "clusters", "list",
		"--format=json(name,location,status,createTime,currentNodeCount,currentMasterVersion)",
	)
}
