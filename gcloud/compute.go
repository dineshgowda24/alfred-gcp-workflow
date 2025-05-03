package gcloud

import (
	"fmt"
	"strings"
)

type ComputeInstance struct {
	Name              string `json:"name"`
	Status            string `json:"status"`
	Zone              string `json:"zone"`
	CPUPlatform       string `json:"cpuPlatform"`
	CreationTimestamp string `json:"creationTimestamp"`
}

func ListComputeInstances(config *Config) ([]ComputeInstance, error) {
	return runGCloudCmd[[]ComputeInstance](config,
		"compute", "instances", "list",
		"--format=json(name,status,zone,cpuPlatform,creationTimestamp)")
}

type ComputeDisk struct {
	Name                   string `json:"name"`
	Status                 string `json:"status"`
	Type                   string `json:"type"`
	SizeGb                 string `json:"sizeGb"`
	Zone                   string `json:"zone"`
	PhysicalBlockSizeBytes string `json:"physicalBlockSizeBytes"`
	CreationTimestamp      string `json:"creationTimestamp"`
}

func (d ComputeDisk) IsRegionSupported(config *Config) bool {
	return isComputeRegionSupported(config)
}

func ListComputeDisks(config *Config) ([]ComputeDisk, error) {
	args := []string{
		"disks", "list",
		"--format=json(name,status,type,sizeGb,zone,physicalBlockSizeBytes,creationTimestamp)",
	}
	if config != nil && config.GetRegionName() != "" {
		args = append(args, fmt.Sprintf("--filter=zone:(%s)", config.GetRegionName()))
	}

	return runGCloudCmd[[]ComputeDisk](config, "compute", args...)
}

type ComputeImage struct {
	Name              string `json:"name"`
	Status            string `json:"status"`
	Architecture      string `json:"architecture"`
	DiskSizeGb        string `json:"diskSizeGb"`
	ArchiveSizeBytes  string `json:"archiveSizeBytes"`
	CreationTimestamp string `json:"creationTimestamp"`
}

func ListComputeImages(config *Config) ([]ComputeImage, error) {
	return runGCloudCmd[[]ComputeImage](
		config, "compute", "images", "list",
		"--format=json(name,status,architecture,diskSizeGb,archiveSizeBytes,creationTimestamp)")
}

type ComputeInstanceTemplate struct {
	CreationTimestamp string `json:"creationTimestamp"`
	Name              string `json:"name"`
	Properties        struct {
		MachineType string `json:"machineType"`
	} `json:"properties"`
}

func (t ComputeInstanceTemplate) IsRegionSupported(config *Config) bool {
	return isComputeRegionSupported(config)
}

func ListComputeInstanceTemplates(config *Config) ([]ComputeInstanceTemplate, error) {
	args := []string{
		"instance-templates", "list",
		"--format=json(creationTimestamp,name,properties.machineType)",
	}

	if config != nil && config.GetRegionName() != "" {
		args = append(args, "--regions="+config.GetRegionName())
	}
	return runGCloudCmd[[]ComputeInstanceTemplate](config, "compute", args...)
}

type ComputeMachineImage struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Status             string `json:"status"`
	TotalStorageBytes  string `json:"totalStorageBytes"`
	CreationTimestamp  string `json:"creationTimestamp"`
	InstanceProperties struct {
		MachineType string `json:"machineType"`
	} `json:"instanceProperties"`
}

func ListComputeMachineImages(config *Config) ([]ComputeMachineImage, error) {
	return runGCloudCmd[[]ComputeMachineImage](
		config,
		"compute", "machine-images", "list",
		"--format=json(name,description,status,totalStorageBytes,creationTimestamp,instanceProperties.machineType)",
	)
}

type ComputeSnapshot struct {
	Name              string `json:"name"`
	Status            string `json:"status"`
	DiskSizeGb        string `json:"diskSizeGb"`
	StorageBytes      string `json:"storageBytes"`
	CreationTimestamp string `json:"creationTimestamp"`
}

func ListComputeSnapshots(config *Config) ([]ComputeSnapshot, error) {
	return runGCloudCmd[[]ComputeSnapshot](
		config,
		"compute", "snapshots", "list",
		"--format=json(name,status,diskSizeGb,storageBytes,creationTimestamp)",
	)
}

func GetAllComputeRegions(config *Config) ([]string, error) {
	type Alias struct {
		Name string `json:"name"`
	}

	regions, err := runGCloudCmd[[]Alias](config, "compute", "regions", "list", "--format=json(name)")
	if err != nil {
		return nil, err
	}

	var regionNames []string
	for _, region := range regions {
		regionNames = append(regionNames, region.Name)
	}

	return regionNames, nil
}

func isComputeRegionSupported(config *Config) bool {
	if config == nil || config.GetRegionName() == "" {
		return true
	}

	regions, err := GetAllComputeRegions(config)
	if err != nil {
		return false
	}

	for _, reg := range regions {
		if strings.EqualFold(reg, config.GetRegionName()) {
			return true
		}
	}

	return false
}
