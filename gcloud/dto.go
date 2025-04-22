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

type Bucket struct {
	CreationTime        string `json:"creation_time"` // Format 2020-09-28T09:24:57+0000 unparsable to time.Time
	DefaultStorageClass string `json:"default_storage_class"`
	Location            string `json:"location"`
	LocationType        string `json:"location_type"`
	Name                string `json:"name"`
	StorageURL          string `json:"storage_url"`
	UpdateTime          string `json:"update_time"` // Format 2020-09-28T09:24:57+0000 unparsable to time.Time
}

type ComputeInstance struct {
	CPUPlatform       string `json:"cpuPlatform"`
	CreationTimestamp string `json:"creationTimestamp"` // Format 2018-11-06T01:59:29.838-08:00
	Name              string `json:"name"`
	Status            string `json:"status"`
	Zone              string `json:"zone"` // Format: projects/{project}/zones/{zone}
}

type ComputeDisk struct {
	CreationTimestamp      string `json:"creationTimestamp"` // Format 2018-11-06T01:59:29.838-08:00
	Name                   string `json:"name"`
	PhysicalBlockSizeBytes string `json:"physicalBlockSizeBytes"`
	SizeGb                 string `json:"sizeGb"`
	Status                 string `json:"status"` // one of [CREATING, RESTORING, FAILED, READY, DELETING]
	Type                   string `json:"type"`   // Format: url/zones/{zone}/diskTypes/{disk-type}
	Zone                   string `json:"zone"`   // Format: url/zones/{zone}
}
