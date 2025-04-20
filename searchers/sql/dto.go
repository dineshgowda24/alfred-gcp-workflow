package sql

import (
	"fmt"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type SQLInstance struct {
	Name                       string
	DatebaseVersion            string
	Region                     string
	InstanceType               string
	Project                    string
	State                      string
	DiskSizeInGB               string
	DatabaseReplicationEnabled bool
	StorageAutoResize          bool
	StorageAutoResizeLimit     string
	Tier                       string
}

func FromGCloudSQLInstance(instance *gcloud.SQLInstance) SQLInstance {
	return SQLInstance{
		Name:                       instance.Name,
		DatebaseVersion:            instance.DatabaseVersion,
		Region:                     instance.Region,
		InstanceType:               instance.InstanceType,
		Project:                    instance.Project,
		State:                      instance.State,
		DiskSizeInGB:               instance.Settings.DataDiskSizeGb,
		DatabaseReplicationEnabled: instance.Settings.DatabaseReplicationEnabled,
		StorageAutoResize:          instance.Settings.StorageAutoResize,
		StorageAutoResizeLimit:     instance.Settings.StorageAutoResizeLimit,
		Tier:                       instance.Settings.Tier,
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
	case "PENDING_CREATE", "UPDATING", "MAINTENANCE":
		icon = "🕒"
	case "FAILED", "DELETING", "SUSPENDED":
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
	return "https://console.cloud.google.com/sql/instances/" + i.Name + "?project=" + config.Project
}
