package memorystore

import (
	"strconv"
	"time"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type RedisInstance struct {
	Name         string
	Region       string
	State        string
	Tier         string
	Memory       int
	RedisVersion string
	ReplicaCount int
	CreatedAt    time.Time
}

func (r RedisInstance) Title() string {
	return r.Name
}

func (r RedisInstance) Subtitle() string {
	if r.State == "READY" {
		return "✅ " + r.RedisVersion + " | " + strconv.Itoa(r.Memory) + " GB | " + "replicas: " + strconv.Itoa(r.ReplicaCount)
	} else if r.State == "CREATING" || r.State == "UPDATING" || r.State == "MAINTENANCE" {
		return "🕒 " + r.RedisVersion + " | " + strconv.Itoa(r.Memory) + " GB | " + "replicas: " + strconv.Itoa(r.ReplicaCount)
	} else {
		return "❌ " + r.RedisVersion + " | " + strconv.Itoa(r.Memory) + " GB | " + "replicas: " + strconv.Itoa(r.ReplicaCount)
	}
}

func (r RedisInstance) URL(config *gcloud.Config) string {
	return "https://console.cloud.google.com/memorystore/redis/instances/details/" + r.Name + "?project=" + config.Project + "&region=" + r.Region
}

func RedisInstanceFromGCloud(instance *gcloud.RedisInstance) RedisInstance {
	return RedisInstance{
		Name:         instance.Name,
		Region:       instance.Region,
		State:        instance.State,
		Tier:         instance.Tier,
		Memory:       instance.Memory,
		RedisVersion: instance.RedisVersion,
		ReplicaCount: instance.ReplicaCount,
		CreatedAt:    instance.CreatedAt,
	}
}
