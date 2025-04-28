package gcloud

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

type ComputeImage struct {
	Architecture      string `json:"architecture"`
	ArchiveSizeBytes  string `json:"archiveSizeBytes"`
	CreationTimestamp string `json:"creationTimestamp"`
	Name              string `json:"name"`
	DiskSizeGb        string `json:"diskSizeGb"`
	Status            string `json:"status"` // one of [FAILED, PENDING, READY]
}

type ComputeInstanceTemplate struct {
	CreationTimestamp string `json:"creationTimestamp"`
	Name              string `json:"name"`
	Properties        struct {
		MachineType string `json:"machineType"`
	} `json:"properties"`
}

// https://cloud.google.com/compute/docs/reference/rest/v1/machineImages/list
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

// https://cloud.google.com/compute/docs/reference/rest/v1/snapshots
type ComputeSnapshot struct {
	Name              string `json:"name"`
	Status            string `json:"status"`
	DiskSizeGb        string `json:"diskSizeGb"`
	StorageBytes      string `json:"storageBytes"`
	CreationTimestamp string `json:"creationTimestamp"`
}

func ListComputeInstances(config *Config) ([]ComputeInstance, error) {
	return runGCloudCmd[[]ComputeInstance](config, "compute", "instances", "list")
}

func ListComputeDisks(config *Config) ([]ComputeDisk, error) {
	return runGCloudCmd[[]ComputeDisk](config, "compute", "disks", "list")
}

func ListComputeImages(config *Config) ([]ComputeImage, error) {
	return runGCloudCmd[[]ComputeImage](config, "compute", "images", "list")
}

func ListComputeInstanceTemplates(config *Config) ([]ComputeInstanceTemplate, error) {
	return runGCloudCmd[[]ComputeInstanceTemplate](config, "compute", "instance-templates", "list")
}

func ListComputeMachineImages(config *Config) ([]ComputeMachineImage, error) {
	return runGCloudCmd[[]ComputeMachineImage](
		config,
		"compute", "machine-images", "list",
		"--format=json(creationTimestamp,description,id,instanceProperties.machineType,name,status,totalStorageBytes)",
	)
}

func ListComputeSnapshots(config *Config) ([]ComputeSnapshot, error) {
	return runGCloudCmd[[]ComputeSnapshot](
		config,
		"compute", "snapshots", "list",
		"--format=json(name,status,diskSizeGb,storageBytes,creationTimestamp)",
	)
}
