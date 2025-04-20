package gcloud

import (
	"encoding/json"
	"os/exec"
)

func ListCloudStorageBuckets(config *Config) ([]Bucket, error) {
	cmd := exec.Command(
		gcloudPath,
		"storage", "buckets", "list",
		"--format=json",
		"--project="+config.Project,
	)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var buckets []Bucket
	err = json.Unmarshal(output, &buckets)
	if err != nil {
		return nil, err
	}
	return buckets, nil
}
