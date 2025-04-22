package compute

import (
	"fmt"
	"log"
	"strconv"
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

type ComputeDisk struct {
	CreationTimestamp time.Time
	Name              string
	SizeGb            int
	BlockSize         int
	Status            string
	Zone              string
	Type              string
}

func FromGCloudComputeDisk(disk *gcloud.ComputeDisk) ComputeDisk {
	creationTime, err := time.Parse("2006-01-02T15:04:05.000-07:00", disk.CreationTimestamp)
	if err != nil {
		log.Println("LOG: compute: Error parsing creation time:", err)
		creationTime = time.Time{}
	}

	var zone string
	zones := strings.Split(disk.Zone, "/")
	if len(zones) > 0 {
		zone = zones[len(zones)-1]
	}

	var diskType string
	diskTypes := strings.Split(disk.Type, "/")
	if len(diskTypes) > 0 {
		diskType = diskTypes[len(diskTypes)-1]
	}

	size, err := strconv.Atoi(disk.SizeGb)
	if err != nil {
		log.Println("LOG: compute: Error parsing size:", err)
		size = 0
	}

	blockSize, err := strconv.Atoi(disk.PhysicalBlockSizeBytes)
	if err != nil {
		log.Println("LOG: compute: Error parsing block size:", err)
		blockSize = 0
	}

	return ComputeDisk{
		CreationTimestamp: creationTime,
		Name:              disk.Name,
		SizeGb:            size,
		Status:            disk.Status,
		Type:              diskType,
		BlockSize:         blockSize,
		Zone:              zone,
	}
}

func (d ComputeDisk) Title() string {
	return d.Name
}

func (d ComputeDisk) Subtitle() string {
	var icon string
	switch d.Status {
	case "READY":
		icon = "🟢"
	case "CREATING", "RESTORING":
		icon = "🕒"
	case "FAILED", "DELETING":
		icon = "❌"
	default:
		icon = "❓"
	}

	return icon + "  " + strconv.Itoa(d.SizeGb) + " GB (" + strconv.Itoa(d.BlockSize) + " blocks) | Created: " + d.CreationTimestamp.Local().Format("Jan 2, 2006 15:04 MST")
}

func (d ComputeDisk) URL(config *gcloud.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/compute/disksDetail/zones/%s/disks/%s?project=%s", d.Zone, d.Name, config.Project)
}
