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

type IAMServiceAccount struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	UniqueID    string `json:"uniqueId"`
}

func ListIAMServiceAccounts(config *Config) ([]IAMServiceAccount, error) {
	return runGCloudCmd[[]IAMServiceAccount](config,
		"iam", "service-accounts", "list", "--format=json(displayName,email,uniqueId)")
}
