package gcloud

type IAMRole struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Title       string `json:"title"`
}

func ListIAMRoles(config *Config) ([]IAMRole, error) {
	return runGCloudCmd[[]IAMRole](config,
		"iam", "roles", "list", "--format=json(description,name,title)")
}
