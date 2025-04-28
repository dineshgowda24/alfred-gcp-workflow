package gcloud

import "time"

// https://cloud.google.com/sql/docs/postgres/admin-api/rest/v1/instances#DatabaseInstance
// https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1/instances#DatabaseInstance
type SQLInstance struct {
	CreateTime      time.Time `json:"createTime"`
	DatabaseVersion string    `json:"databaseVersion"`
	InstanceType    string    `json:"instanceType"`
	Name            string    `json:"name"`
	State           string    `json:"state"`
	Settings        struct {
		DataDiskSizeGb string `json:"dataDiskSizeGb"`
		Tier           string `json:"tier"`
	} `json:"settings"`
}

func ListSQLInstances(config *Config) ([]SQLInstance, error) {
	return runGCloudCmd[[]SQLInstance](
		config,
		"sql", "instances", "list",
		"--format=json(name,databaseVersion,instanceType,project,state,settings.dataDiskSizeGb,settings.tier)",
	)
}
