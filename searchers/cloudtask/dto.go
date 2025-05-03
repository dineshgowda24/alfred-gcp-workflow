package cloudtask

import (
	"fmt"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Queue struct {
	Name       string
	Region     string
	RateLimits struct {
		MaxBurstSize            int
		MaxConcurrentDispatches int
		MaxDispatchesPerSecond  float64
	}
}

func (q Queue) Title() string {
	return q.Name
}

func (q Queue) Subtitle() string {
	return fmt.Sprintf("📍 %s | 🚀 %g/s, burst %d, max %d",
		q.Region,
		q.RateLimits.MaxDispatchesPerSecond,
		q.RateLimits.MaxBurstSize,
		q.RateLimits.MaxConcurrentDispatches,
	)
}

func (q Queue) URL(config *gcloud.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/cloudtasks/queue/%s/%s/tasks?project=%s",
		q.Region,
		q.Name,
		config.Project,
	)
}

func FromGCloudCloudTaskQueue(gctq *gcloud.CloudTaskQueue) Queue {
	parts := strings.Split(gctq.Name, "/")

	var region string
	var name string
	for i, part := range parts {
		switch part {
		case "locations":
			if i+1 < len(parts) {
				region = parts[i+1]
			}
		case "queues":
			if i+1 < len(parts) {
				name = parts[i+1]
			}
		}
	}

	return Queue{
		Name:   name,
		Region: region,
		RateLimits: struct {
			MaxBurstSize            int
			MaxConcurrentDispatches int
			MaxDispatchesPerSecond  float64
		}{
			MaxBurstSize:            gctq.RateLimits.MaxBurstSize,
			MaxConcurrentDispatches: gctq.RateLimits.MaxConcurrentDispatches,
			MaxDispatchesPerSecond:  gctq.RateLimits.MaxDispatchesPerSecond,
		},
	}
}
