package gcloud

// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances
type RedisInstance struct {
	FullName     string `json:"name"`
	DisplayName  string `json:"displayName"`
	LocationId   string `json:"currentLocationId"`
	State        string `json:"state"`
	Memory       int    `json:"memorySizeGb"`
	RedisVersion string `json:"redisVersion"`
	ReplicaCount int    `json:"replicaCount"`
}

func ListRedisInstances(config *Config) ([]RedisInstance, error) {
	return runGCloudCmd[[]RedisInstance](
		config,
		"redis", "instances", "list",
		"--region=-",
		"--format=json(name,displayName,currentLocationId,state,memorySizeGb,redisVersion,replicaCount)",
	)
}
