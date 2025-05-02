package gcloud

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strings"
)

type GCloudInfo struct {
	Config struct {
		Paths struct {
			GlobalConfigDir string `json:"global_config_dir"`
		} `json:"paths"`
	} `json:"config"`
	Installation struct {
		Components struct {
			Beta                string `json:"beta"`
			BQ                  string `json:"bq"`
			Core                string `json:"core"`
			GCloudCRC32C        string `json:"gcloud-crc32c"`
			GKEGCloudAuthPlugin string `json:"gke-gcloud-auth-plugin"`
			Gsutil              string `json:"gsutil"`
			Kubectl             string `json:"kubectl"`
		} `json:"components"`
	} `json:"installation"`
}

func GetGCloudInfo(cliPath string) (*GCloudInfo, error) {
	cmd := exec.Command(cliPath, "info", "--format=json(config,installation)")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var info GCloudInfo
	if err := json.Unmarshal(out, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func NormalizeGCloudPath(path string) string {
	if path == "" {
		return ""
	}
	if !strings.HasSuffix(path, "gcloud") {
		return filepath.Join(path, "gcloud")
	}
	return path
}
