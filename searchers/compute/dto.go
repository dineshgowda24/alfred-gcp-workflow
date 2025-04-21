package compute

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type ComputeInstance struct {
	CPUPlatform       string
	CreationTimestamp time.Time
	Name              string
	Status            string
	Zone              string
}

func FromGCloudComputeInstance(instance *gcloud.ComputeInstance) ComputeInstance {
	creationTime, err := time.Parse("2006-01-02T15:04:05.000-07:00", instance.CreationTimestamp)
	if err != nil {
		log.Println("LOG: compute: Error parsing creation time:", err)
		creationTime = time.Time{}
	}

	var zone string
	zones := strings.Split(instance.Zone, "/")
	if len(zone) > 0 {
		zone = zones[len(zones)-1]
	}

	return ComputeInstance{
		CPUPlatform:       instance.CPUPlatform,
		CreationTimestamp: creationTime,
		Name:              instance.Name,
		Status:            instance.Status,
		Zone:              zone,
	}
}

func (i ComputeInstance) Title() string {
	return i.Name
}

func (i ComputeInstance) Subtitle() string {
	var icon string
	switch i.Status {
	case "RUNNING":
		icon = "🟢"
	case "PROVISIONING", "STOPPING", "SUSPENDED":
		icon = "🕒"
	case "TERMINATED", "REPAIRING":
		icon = "❌"
	default:
		icon = "❓"
	}

	return icon + "  CPU: " + i.CPUPlatform + " | Created: " + i.CreationTimestamp.Local().Format("Jan 2, 2006 15:04 MST")
}

func (i ComputeInstance) URL(config *gcloud.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/compute/instancesDetail/zones/%s/instances/%s?project=%s", i.Zone, i.Name, config.Project)
}
