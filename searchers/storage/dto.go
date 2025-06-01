package storage

import (
	"time"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/util"
)

type Bucket struct {
	CreationTime        time.Time
	DefaultStorageClass string
	Location            string
	LocationType        string
	Name                string
	StorageURL          string
	UpdateTime          time.Time
}

func (b Bucket) Title() string {
	return b.Name
}

func (b Bucket) Subtitle() string {
	return b.LocationType + " | " + b.DefaultStorageClass + " | LastUpdated: " + util.FormatLocalTime(b.UpdateTime)
}

func (b Bucket) URL(config *gcloud.Config) string {
	return "https://console.cloud.google.com/storage/browser/" + b.Name + "?project=" + config.Project
}

func FromGCloudStorageBucket(bucket *gcloud.Bucket) Bucket {
	updateTime, err := time.Parse("2006-01-02T15:04:05-0700", bucket.UpdateTime)
	if err != nil {
		updateTime = time.Time{}
	}

	creationTime, err := time.Parse("2006-01-02T15:04:05-0700", bucket.CreationTime)
	if err != nil {
		creationTime = time.Time{}
	}

	return Bucket{
		CreationTime:        creationTime,
		DefaultStorageClass: bucket.DefaultStorageClass,
		Location:            bucket.Location,
		LocationType:        bucket.LocationType,
		Name:                bucket.Name,
		StorageURL:          bucket.StorageURL,
		UpdateTime:          updateTime,
	}
}
