package gcloud

// https://cloud.google.com/filestore/docs/reference/rest/v1/projects.locations.instances
type FilestoreInstance struct {
	Name       string `json:"name"`
	State      string `json:"state"`
	Tier       string `json:"tier"`
	FileShares []struct {
		CapacityGb string `json:"capacityGb"`
		Name       string `json:"name"`
	} `json:"fileShares"`
	CreatedAt string `json:"createTime"`
}

func ListFilestoreInstances(config *Config) ([]FilestoreInstance, error) {
	return runGCloudCmd[[]FilestoreInstance](
		config,
		"filestore", "instances", "list",
		"--format=json(name,state,tier,createTime,fileShares)",
	)
}
