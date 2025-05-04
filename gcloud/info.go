package gcloud

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const execPerm = 0o111

var (
	ErrGCloudNotFound = errors.New("gcloud binary not found")
	ErrNonExecutable  = errors.New("provided path is not executable")
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

func GetGCloudInfo(binPath string) (*GCloudInfo, error) {
	stat, err := os.Stat(binPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrGCloudNotFound
		}
		return nil, err
	}

	if stat.IsDir() || stat.Mode().Perm()&execPerm == 0 {
		return nil, ErrNonExecutable
	}

	out, err := exec.Command(binPath, "info", "--format=json(config,installation)").Output()
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
	normalized := strings.TrimSpace(path)
	if normalized == "" {
		return ""
	}

	if strings.HasPrefix(normalized, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Println("error getting home directory:", err)
			return normalized
		}

		normalized = filepath.Join(homeDir, normalized[1:])
	}

	normalized = filepath.Clean(normalized)
	if filepath.Base(normalized) != "gcloud" {
		return filepath.Join(normalized, "gcloud")
	}

	return normalized
}
