package storage

import (
	"log"
	"time"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
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
	localTime := b.UpdateTime.Local()
	return b.LocationType + " Region | " + b.DefaultStorageClass + " | LastUpdated: " + localTime.Format("Jan 2, 2006 15:04 MST")
}

func (b Bucket) URL(config *gcloud.Config) string {
	return "https://console.cloud.google.com/storage/browser/" + b.Name + "?project=" + config.Project
}

func FromGCloudStorageBucket(bucket *gcloud.Bucket) Bucket {
	updateTime, err := time.Parse("2006-01-02T15:04:05-0700", bucket.UpdateTime)
	if err != nil {
		log.Println("LOG: Error parsing update time:", err)
		updateTime = time.Time{}
	}

	creationTime, err := time.Parse("2006-01-02T15:04:05-0700", bucket.CreationTime)
	if err != nil {
		log.Println("LOG: Error parsing creation time:", err)
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
