package gcloud

import (
	"encoding/json"
	"os/exec"
)

func ListRedisInstances(config *Config) ([]RedisInstance, error) {
	cmd := exec.Command(
		gcloudPath,
		"redis", "instances", "list",
		"--format=json",
		"--project="+config.Project,
		"--region=-", // Needed to avoid region error, or use loop for all regions
	)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var instances []RedisInstance
	err = json.Unmarshal(output, &instances)
	if err != nil {
		return nil, err
	}
	return instances, nil
}
