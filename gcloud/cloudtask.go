package gcloud

type CloudTaskQueue struct {
	Name       string `json:"name"`
	RateLimits struct {
		MaxBurstSize            int     `json:"maxBurstSize"`
		MaxConcurrentDispatches int     `json:"maxConcurrentDispatches"`
		MaxDispatchesPerSecond  float64 `json:"maxDispatchesPerSecond"`
	} `json:"rateLimits"`
}

func (q CloudTaskQueue) IsRegionSupported(config *Config) bool {
	if config == nil || config.Region == nil || config.Region.Name == "" {
		return true
	}

	type Location struct {
		LocationID string `json:"locationId"`
	}

	locations, err := runGCloudCmd[[]Location](config, "tasks", "locations", "list", "--format=json(locationId)")
	if err != nil {
		return false
	}

	for _, loc := range locations {
		if loc.LocationID == config.Region.Name {
			return true
		}
	}

	return false
}

func ListCloudTaskQueues(config *Config) ([]CloudTaskQueue, error) {
	args := []string{"queues", "list", "--format=json(name,rateLimits)"}

	if config != nil && config.Region != nil {
		args = append(args, "--location="+config.Region.Name)
	}

	return runGCloudCmd[[]CloudTaskQueue](config, "tasks", args...)
}
