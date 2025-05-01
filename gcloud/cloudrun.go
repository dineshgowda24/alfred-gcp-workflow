package gcloud

// https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services
type CloudRunService struct {
	Metadata struct {
		Name   string `json:"name"`
		Labels struct {
			Location string `json:"cloud.googleapis.com/location"`
		} `json:"labels"`
		CreationTimestamp string `json:"creationTimestamp"`
	}
}

func ListCloudRunServices(config *Config) ([]CloudRunService, error) {
	return runGCloudCmd[[]CloudRunService](
		config,
		"run", "services", "list",
		"--format=json(metadata.name,metadata.labels,metadata.creationTimestamp)",
	)
}

// https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services
// Since we are using --v2 flag while listing, we can refer the same services reference
type CloudRunFunction struct {
	Name        string `json:"name"`
	State       string `json:"state"`
	Environment string `json:"environment"`
	UpdateTime  string `json:"updateTime"`
	BuildConfig struct {
		Runtime string `json:"runtime"`
	} `json:"buildConfig"`
}

// ListCloudRunFunctions retrieves only Cloud Functions (Gen 1).
// By default, `gcloud functions list --v2` includes both Gen 1 and Gen 2,
// so we apply a filter on `environment=GEN_1` to exclude Gen 2 services.
// The --v2 flag ensures a consistent data format across environments.
func ListCloudRunFunctions(config *Config) ([]CloudRunFunction, error) {
	return runGCloudCmd[[]CloudRunFunction](
		config,
		"functions", "list", "--v2",
		"--format=json(name,state,environment,updateTime,buildConfig.runtime)",
		"--filter=environment=GEN_1",
	)
}
