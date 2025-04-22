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
	CreationTimestamp string `json:"creationTimestamp"` // Format 2018-11-06T01:59:29.838-08:00
	Name              string `json:"name"`
	DiskSizeGb        string `json:"diskSizeGb"`
	Status            string `json:"status"` // one of [FAILED, PENDING, READY]
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
