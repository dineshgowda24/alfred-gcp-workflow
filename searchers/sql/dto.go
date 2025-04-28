package sql

import (
	"fmt"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type SQLInstance struct {
	Name            string
	DatebaseVersion string
	InstanceType    string
	State           string
	DiskSizeInGB    string
	Tier            string
}

func FromGCloudSQLInstance(instance *gcloud.SQLInstance) SQLInstance {
	return SQLInstance{
		Name:            instance.Name,
		DatebaseVersion: instance.DatabaseVersion,
		InstanceType:    instance.InstanceType,
		State:           instance.State,
		DiskSizeInGB:    instance.Settings.DataDiskSizeGb,
		Tier:            instance.Settings.Tier,
	}
}

func (i SQLInstance) Title() string {
	return i.Name
}

func (i SQLInstance) Subtitle() string {
	var icon string
	switch i.State {
	case "RUNNABLE":
		icon = "🟢"
	case "PENDING_CREATE", "MAINTENANCE", "ONLINE_MAINTENANCE":
		icon = "🕒"
	case "PENDING_DELETE", "FAILED", "SUSPENDED":
		icon = "❌"
	default:
		icon = "❓"
	}

	var role string
	switch i.InstanceType {
	case "READ_REPLICA_INSTANCE":
		role = "Replica"
	case "CLOUD_SQL_INSTANCE":
		role = "Master"
	default:
		role = "Unknown"
	}

	version := strings.TrimPrefix(i.DatebaseVersion, "POSTGRES_")
	version = strings.ReplaceAll(version, "_X", ".x")
	version = strings.ReplaceAll(version, "_", ".")

	return fmt.Sprintf("%s %s | version: %s | %s GB | %s", icon, role, version, i.DiskSizeInGB, i.Tier)
}

func (i SQLInstance) URL(config *gcloud.Config) string {
	return fmt.Sprintf(
		"https://console.cloud.google.com/sql/instances/%s?project=%s",
		i.Name, config.Project)
}
