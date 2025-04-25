package env

import (
	"os"
	"strings"
)

const (
	EnvGCloudCliPath = "ALFRED_GCP_WORKFLOW_GCLOUD_PATH"
)

func GCloudCliPath() string {
	val := os.Getenv(EnvGCloudCliPath)
	return strings.TrimSpace(val)
}
