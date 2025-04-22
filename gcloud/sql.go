package gcloud

import "time"

type SQLInstance struct {
	CreateTime               time.Time `json:"createTime"`
	DatabaseInstalledVersion string    `json:"databaseInstalledVersion"`
	DatabaseVersion          string    `json:"databaseVersion"`
	GCEZone                  string    `json:"gceZone"`
	InstanceType             string    `json:"instanceType"`
	Name                     string    `json:"name"`
	Project                  string    `json:"project"`
	Region                   string    `json:"region"`
	State                    string    `json:"state"`
	Settings                 struct {
		DataDiskSizeGb             string `json:"dataDiskSizeGb"`
		DataDiskType               string `json:"dataDiskType"`
		DatabaseReplicationEnabled bool   `json:"databaseReplicationEnabled"`
		StorageAutoResize          bool   `json:"storageAutoResize"`
		StorageAutoResizeLimit     string `json:"storageAutoResizeLimit"`
		Tier                       string `json:"tier"`
	} `json:"settings"`
}

func ListSQLInstances(config *Config) ([]SQLInstance, error) {
	return runGCloudCmd[[]SQLInstance](config, "sql", "instances", "list")
}
