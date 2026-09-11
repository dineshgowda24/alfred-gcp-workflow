package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	CacheTTLEnv = "ALFRED_GCP_WORKFLOW_CACHE_TTL_SECONDS"
	AuthUserEnv = "ALFRED_GCP_WORKFLOW_AUTHUSER"
)

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

// AuthUser returns the configured authuser value for console URLs.
// See gcloud.AuthUser for how the value is interpreted.
func AuthUser() string {
	return strings.TrimSpace(os.Getenv(AuthUserEnv))
}
