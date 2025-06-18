package gcloud

type Dashboard struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
}

func ListMonitoringDashboards(config *Config) ([]Dashboard, error) {
	return runGCloudCmd[[]Dashboard](config,
	 "monitoring", "dashboards", "list", "--format=json(displayName,name)")
}
