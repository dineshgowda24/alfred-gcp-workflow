package gcloud

import (
	"encoding/json"
	"os/exec"
)

func ListSQLInstances(config *Config) ([]SQLInstance, error) {
	cmd := exec.Command(gcloudPath, "sql", "instances", "list", "--format=json", "--project="+config.Project)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var instances []SQLInstance
	err = json.Unmarshal(output, &instances)
	if err != nil {
		return nil, err
	}

	return instances, nil
}
