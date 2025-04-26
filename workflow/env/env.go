package env

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	EnvGCloudCliPath   = "ALFRED_GCP_WORKFLOW_GCLOUD_PATH"
	EnvCacheTTLSeconds = "ALFRED_GCP_WORKFLOW_CACHE_TTL_SECONDS"
)

func GCloudCliPath() string {
	val := os.Getenv(EnvGCloudCliPath)
	return strings.TrimSpace(val)
}

func CacheTTLDuration(defaultTTL time.Duration) time.Duration {
	val := os.Getenv(EnvCacheTTLSeconds)
	if val == "" {
		log.Println("LOG: Cache TTL not set, using default value:", defaultTTL)
		return defaultTTL
	}

	parsed, err := strconv.Atoi(val)
	if err != nil {
		log.Println("LOG: Error parsing cache TTL, using default value:", err)
		return defaultTTL
	}
	return time.Duration(parsed) * time.Second
}
