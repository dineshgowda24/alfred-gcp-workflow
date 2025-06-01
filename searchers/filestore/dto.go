package filestore

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/util"
)

type Instance struct {
	FullName        string
	Id              string
	Location        string
	State           string
	Tier            string
	TotalCapacityGb int
	CreatedAt       time.Time
}

func (i Instance) Title() string {
	return i.Id
}

func (i Instance) Subtitle() string {
	var icon string
	switch i.State {
	case "READY":
		icon = "🟢"
	case "CREATING", "REPAIRING", "RESTORING", "RESUMING", "REVERTING", "PROMOTING":
		icon = "🕒"
	case "DELETING", "SUSPENDING", "SUSPENDED", "ERROR":
		icon = "❌"
	default:
		icon = "❓"
	}
	return fmt.Sprintf("%s %s | %d GB | %s",
		icon, i.Tier, i.TotalCapacityGb, util.FormatTime(i.CreatedAt),
	)
}

func (i Instance) URL(config *gcloud.Config) string {
	return fmt.Sprintf(
		"https://console.cloud.google.com/filestore/instances/locations/%s/id/%s?project=%s",
		i.Location,
		i.Id,
		config.Project,
	)
}

func FromGCloudInstance(instance *gcloud.FilestoreInstance) Instance {
	createdAt, err := time.Parse(time.RFC3339, instance.CreatedAt)
	if err != nil {
		createdAt = time.Time{}
	}

	var totalCapacityGb int
	if len(instance.FileShares) > 0 {
		capacity, err := strconv.Atoi(instance.FileShares[0].CapacityGb)
		if err != nil {
			capacity = 0
		}

		totalCapacityGb = capacity
	}

	var location, id string
	parts := strings.Split(instance.Name, "/")
	for i := 0; i < len(parts)-1; i++ {
		switch parts[i] {
		case "locations":
			location = parts[i+1]
		case "instances":
			id = parts[i+1]
		}
	}

	return Instance{
		Id:              id,
		FullName:        instance.Name,
		State:           instance.State,
		Tier:            instance.Tier,
		Location:        location,
		TotalCapacityGb: totalCapacityGb,
		CreatedAt:       createdAt,
	}
}
