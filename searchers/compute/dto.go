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

type ComputeImage struct {
	Architecture      string
	ArchiveSizeBytes  int
	CreationTimestamp time.Time
	Name              string
	DiskSizeGb        int
	Status            string
}

func FromGCloudComputeImage(image *gcloud.ComputeImage) ComputeImage {
	creationTime, err := time.Parse("2006-01-02T15:04:05.000-07:00", image.CreationTimestamp)
	if err != nil {
		log.Println("LOG: compute: Error parsing creation time:", err)
		creationTime = time.Time{}
	}

	size, err := strconv.Atoi(image.DiskSizeGb)
	if err != nil {
		log.Println("LOG: compute: Error parsing size:", err)
		size = 0
	}

	archiveSize, err := strconv.Atoi(image.ArchiveSizeBytes)
	if err != nil {
		log.Println("LOG: compute: Error parsing archive size:", err)
		archiveSize = 0
	}

	return ComputeImage{
		Architecture:      image.Architecture,
		ArchiveSizeBytes:  archiveSize,
		Name:              image.Name,
		DiskSizeGb:        size,
		Status:            image.Status,
		CreationTimestamp: creationTime,
	}
}

func (i ComputeImage) Title() string {
	return i.Name
}

func (i ComputeImage) Subtitle() string {
	var icon string
	switch i.Status {
	case "READY":
		icon = "🟢"
	case "PENDING":
		icon = "🕒"
	case "FAILED":
		icon = "❌"
	default:
		icon = "❓"
	}

	return icon + "  " + strconv.Itoa(i.DiskSizeGb) + " GB | Created: " + i.CreationTimestamp.Local().Format("Jan 2, 2006 15:04 MST")
}

func (i ComputeImage) URL(config *gcloud.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/compute/imagesDetail/projects/%s/global/images/%s?project=%s", config.Project, i.Name, config.Project)
}

type ComputeInstanceTemplate struct {
	CreationTimestamp time.Time
	Name              string
	MachineType       string
}

func FromGCloudComputeInstanceTemplate(template *gcloud.ComputeInstanceTemplate) ComputeInstanceTemplate {
	creationTime, err := time.Parse("2006-01-02T15:04:05.000-07:00", template.CreationTimestamp)
	if err != nil {
		log.Println("LOG: compute: Error parsing creation time:", err)
		creationTime = time.Time{}
	}

	return ComputeInstanceTemplate{
		CreationTimestamp: creationTime,
		Name:              template.Name,
		MachineType:       template.Properties.MachineType,
	}
}

func (t ComputeInstanceTemplate) Title() string {
	return t.Name
}

func (t ComputeInstanceTemplate) Subtitle() string {
	return t.MachineType + " | Created: " + t.CreationTimestamp.Local().Format("Jan 2, 2006 15:04 MST")
}

func (t ComputeInstanceTemplate) URL(config *gcloud.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/compute/instanceTemplates/details/%s?project=%s", t.Name, config.Project)
}

type MachineImage struct {
	Name              string
	Description       string
	Status            string
	TotalStorage      string
	CreationTimestamp time.Time
	MachineType       string
}

func (m MachineImage) Title() string {
	return m.Name
}

func (m MachineImage) Subtitle() string {
	var icon string
	switch m.Status {
	case "READY":
		icon = "🟢"
	case "CREATING", "UPLOADING":
		icon = "🕒"
	case "INVALID", "DELETING":
		icon = "❌"
	default:
		icon = "❓"
	}

	return fmt.Sprintf("%s %s | %s | Created: %s",
		icon, m.MachineType, m.TotalStorage, m.CreationTimestamp.Local().Format("Jan 2, 2006 15:04 MST"))
}

func (m MachineImage) URL(config *gcloud.Config) string {
	return fmt.Sprintf(
		"https://console.cloud.google.com/compute/machineImages/details/%s?project=%s",
		m.Name, config.Project)
}

func FromGCloudComputeMachineImage(image *gcloud.ComputeMachineImage) MachineImage {
	creationTime, err := time.Parse(time.RFC3339, image.CreationTimestamp)
	if err != nil {
		log.Println("LOG: compute: Error parsing creation time:", err)
		creationTime = time.Time{}
	}

	size, err := strconv.Atoi(image.TotalStorageBytes)
	if err != nil {
		log.Println("LOG: compute: Error parsing size:", err)
		size = 0
	}

	storageIdentifiers := []string{"B", "KB", "MB", "GB", "TB"}
	idx := 0
	for size > 1024 && idx < len(storageIdentifiers)-1 {
		size /= 1024
		idx++
	}

	return MachineImage{
		Name:              image.Name,
		Description:       image.Description,
		Status:            image.Status,
		TotalStorage:      fmt.Sprintf("%d %s", size, storageIdentifiers[idx]),
		MachineType:       image.InstanceProperties.MachineType,
		CreationTimestamp: creationTime,
	}
}

type Snapshot struct {
	Name              string
	Status            string
	DiskSizeGb        int
	StorageSize       string
	CreationTimestamp time.Time
}

func (s Snapshot) Title() string {
	return s.Name
}

func (s Snapshot) Subtitle() string {
	var icon string
	switch s.Status {
	case "READY":
		icon = "🟢"
	case "CREATING", "UPLOADING":
		icon = "🕒"
	case "FAILED", "DELETING":
		icon = "❌"
	default:
		icon = "❓"
	}

	return fmt.Sprintf("%s Size: %s | Disk Size: %d GB | Created: %s",
		icon, s.StorageSize, s.DiskSizeGb, s.CreationTimestamp.Local().Format("Jan 2, 2006 15:04 MST"))
}

func (s Snapshot) URL(config *gcloud.Config) string {
	return fmt.Sprintf(
		"https://console.cloud.google.com/compute/snapshotsDetail/projects/%s/global/snapshots/%s?project=%s",
		config.Project, s.Name, config.Project)
}

func FromGCloudComputeSnapshot(snapshot *gcloud.ComputeSnapshot) Snapshot {
	creationTime, err := time.Parse(time.RFC3339, snapshot.CreationTimestamp)
	if err != nil {
		log.Println("error parsing creation time:", err)
		creationTime = time.Time{}
	}

	size, err := strconv.Atoi(snapshot.DiskSizeGb)
	if err != nil {
		log.Println("error parsing disk size:", err)
		size = 0
	}

	storageSize, err := strconv.Atoi(snapshot.StorageBytes)
	if err != nil {
		log.Println("LOG: compute: Error parsing storage size:", err)
		storageSize = 0
	}

	storageIdentifiers := []string{"B", "KB", "MB", "GB", "TB"}
	idx := 0
	for storageSize > 1024 && idx < len(storageIdentifiers)-1 {
		storageSize /= 1024
		idx++
	}

	return Snapshot{
		Name:              snapshot.Name,
		Status:            snapshot.Status,
		DiskSizeGb:        size,
		StorageSize:       fmt.Sprintf("%d %s", storageSize, storageIdentifiers[idx]),
		CreationTimestamp: creationTime.Local(),
	}
}
