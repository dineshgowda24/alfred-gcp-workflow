package gcloud

import "time"

type Config struct {
	Name    string
	Project string
}

type GCloudConfig struct {
	Name       string `json:"name"`
	Properties struct {
		Core struct {
			Project string `json:"project"`
		} `json:"core"`
	} `json:"properties"`
}

func (g *GCloudConfig) ToConfig() *Config {
	return &Config{
		Name:    g.Name,
		Project: g.Properties.Core.Project,
	}
}

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

type RedisInstance struct {
	CreatedAt    time.Time `json:"createTime"`
	Name         string    `json:"displayName"`
	Region       string    `json:"currentLocationId"`
	State        string    `json:"state"`
	Tier         string    `json:"tier"`
	Memory       int       `json:"memorySizeGb"`
	RedisVersion string    `json:"redisVersion"`
	ReplicaCount int       `json:"replicaCount"`
}

type PubSubTopic struct {
	Name string `json:"name"` // Format: projects/{project}/topics/{topic-id}
}

type PubSubSubscription struct {
	AckDeadlineSeconds       int    `json:"ackDeadlineSeconds"`
	MessageRetentionDuration string `json:"messageRetentionDuration"`
	Name                     string `json:"name"` // Format: projects/{project}/subscriptions/{subscription-id}
	State                    string `json:"state"`
	Topic                    string `json:"topic"` // Format: projects/{project}/topics/{topic-id}
}
