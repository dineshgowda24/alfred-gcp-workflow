package env

import (
	"os"
	"strconv"
	"time"
)

const CacheTTLEnv = "ALFRED_GCP_WORKFLOW_CACHE_TTL_SECONDS"

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
