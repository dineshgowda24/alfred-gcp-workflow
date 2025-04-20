package memorystore

import (
	"strconv"
	"strings"
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
	version := strings.TrimPrefix(r.RedisVersion, "REDIS_")
	version = strings.ReplaceAll(version, "_X", ".x")
	version = strings.ReplaceAll(version, "_", ".")

	info := "version: " + version + " | " + strconv.Itoa(r.Memory) + " GB | replicas: " + strconv.Itoa(r.ReplicaCount)

	switch r.State {
	case "READY":
		return "🟢 " + info
	case "CREATING", "UPDATING", "MAINTENANCE":
		return "🕒 " + info
	default:
		return "❌ " + info
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
