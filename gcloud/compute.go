package gcloud

import (
	"encoding/json"
	"os/exec"
)

func ListComputeInstances(config *Config) ([]ComputeInstance, error) {
	cmd := exec.Command(gcloudPath, "compute", "instances", "list", "--format=json", "--configuration="+config.Name)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var instances []ComputeInstance
	err = json.Unmarshal(output, &instances)
	if err != nil {
		return nil, err
	}

	return instances, nil
}

func ListComputeDisks(config *Config) ([]ComputeDisk, error) {
	cmd := exec.Command(gcloudPath, "compute", "disks", "list", "--format=json", "--configuration="+config.Name)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var disks []ComputeDisk
	err = json.Unmarshal(output, &disks)
	if err != nil {
		return nil, err
	}

	return disks, nil
}
