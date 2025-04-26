package env

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	GCloudCliPathEnv = "ALFRED_GCP_WORKFLOW_GCLOUD_PATH"
	CacheTTLEnv      = "ALFRED_GCP_WORKFLOW_CACHE_TTL_SECONDS"
)

func GCloudCliPath() string {
	val := os.Getenv(GCloudCliPathEnv)
	return strings.TrimSpace(val)
}

func CacheTTLDuration(defaultTTL time.Duration) time.Duration {
	val := os.Getenv(CacheTTLEnv)
	if val == "" {
		return defaultTTL
	}

	ttl, err := strconv.Atoi(val)
	if err != nil {
		return defaultTTL
	}
	return time.Second * time.Duration(ttl)
}
