package gcloud

type Bucket struct {
	CreationTime        string `json:"creation_time"` // Format 2020-09-28T09:24:57+0000 unparsable to time.Time
	DefaultStorageClass string `json:"default_storage_class"`
	Location            string `json:"location"`
	LocationType        string `json:"location_type"`
	Name                string `json:"name"`
	StorageURL          string `json:"storage_url"`
	UpdateTime          string `json:"update_time"` // Format 2020-09-28T09:24:57+0000 unparsable to time.Time
}

func ListCloudStorageBuckets(config *Config) ([]Bucket, error) {
	return runGCloudCmd[[]Bucket](config, "storage", "buckets", "list")
}
