package memorystore

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type RedisInstance struct {
	Id           string
	Name         string
	FullName     string
	LocationId   string
	Region       string
	State        string
	Memory       int
	RedisVersion string
	ReplicaCount int
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
	case "DELETING", "REPAIRING", "IMPORTING", "FAILING_OVER":
		return "❌ " + info
	default:
		return "❓ " + info
	}
}

func (r RedisInstance) URL(config *gcloud.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/memorystore/redis/locations/%s/instances/%s/details/overview?project=%s",
		r.Region, r.Id, config.Project)
}

func RedisInstanceFromGCloud(instance *gcloud.RedisInstance) RedisInstance {
	var region string
	var id string
	words := strings.Split(instance.FullName, "/")
	for i, word := range words {
		if word == "locations" && i+1 < len(words) {
			region = words[i+1]
		}
		if word == "instances" && i+1 < len(words) {
			id = words[i+1]
		}
	}

	log.Println("RedisInstanceFromGCloud", words)

	return RedisInstance{
		Id:           id,
		Name:         instance.DisplayName,
		FullName:     instance.FullName,
		LocationId:   instance.LocationId,
		Region:       region,
		State:        instance.State,
		Memory:       instance.Memory,
		RedisVersion: instance.RedisVersion,
		ReplicaCount: instance.ReplicaCount,
	}
}
