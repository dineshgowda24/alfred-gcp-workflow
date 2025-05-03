package gcloud

import (
	"strings"
)

type RedisInstance struct {
	FullName     string `json:"name"`
	DisplayName  string `json:"displayName"`
	LocationId   string `json:"currentLocationId"`
	State        string `json:"state"`
	Memory       int    `json:"memorySizeGb"`
	RedisVersion string `json:"redisVersion"`
	ReplicaCount int    `json:"replicaCount"`
}

func (r RedisInstance) IsRegionSupported(config *Config) bool {
	if config == nil || config.GetRegionName() == "" {
		return true
	}

	type Alias struct {
		LocationId string `json:"locationId"`
	}

	regions, err := runGCloudCmd[[]Alias](config, "redis", "regions", "list", "--format=json(locationId)")
	if err != nil {
		return false
	}

	for _, r := range regions {
		if strings.EqualFold(r.LocationId, config.GetRegionName()) {
			return true
		}
	}

	return false
}

func ListRedisInstances(config *Config) ([]RedisInstance, error) {
	args := []string{
		"instances", "list",
		"--format=json(name,displayName,currentLocationId,state,memorySizeGb,redisVersion,replicaCount)",
	}

	if config != nil && config.GetRegionName() != "" {
		args = append(args, "--region="+config.GetRegionName())
	} else {
		args = append(args, "--region=-")
	}

	return runGCloudCmd[[]RedisInstance](config, "redis", args...)
}
