package gcloud

import "time"

type RedisInstance struct {
	CreatedAt    time.Time `json:"createTime"`
	Name         string    `json:"displayName"`
	Region       string    `json:"currentLocationId"`
	State        string    `json:"state"`
	Tier         string    `json:"tier"`
	Memory       int       `json:"memorySizeGb"`
	RedisVersion string    `json:"redisVersion"`
	ReplicaCount int       `json:"replicaCount"`
}

func ListRedisInstances(config *Config) ([]RedisInstance, error) {
	return runGCloudCmd[[]RedisInstance](config, "redis", "instances", "list", "--region=-")
}
