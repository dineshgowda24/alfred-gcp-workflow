package gcloud

type Bucket struct {
	CreationTime        string `json:"creation_time"`
	DefaultStorageClass string `json:"default_storage_class"`
	Location            string `json:"location"`
	LocationType        string `json:"location_type"`
	Name                string `json:"name"`
	StorageURL          string `json:"storage_url"`
	UpdateTime          string `json:"update_time"`
}

func ListCloudStorageBuckets(config *Config) ([]Bucket, error) {
	return runGCloudCmd[[]Bucket](config, "storage", "buckets", "list")
}
