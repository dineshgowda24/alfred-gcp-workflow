package cloudrun

import (
	"fmt"
	"strings"
	"time"

	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/util"
)

type Service struct {
	Name         string
	Location     string
	CreationTime time.Time
}

func (s Service) Title() string {
	return s.Name
}

func (s Service) Subtitle() string {
	return fmt.Sprintf("Location: %s | Created: %s", s.Location, util.FormatTime(s.CreationTime))
}

func (s Service) URL(config *gc.Config) string {
	return fmt.Sprintf(
		"https://console.cloud.google.com/run/detail/%s/%s/metrics?project=%s",
		s.Location,
		s.Name,
		config.Project,
	)
}

func FromGCloudCloudRunService(gcs *gc.CloudRunService) Service {
	createTime, err := time.Parse(time.RFC3339, gcs.Metadata.CreationTimestamp)
	if err != nil {
		createTime = time.Time{}
	}

	return Service{
		Name:         gcs.Metadata.Name,
		Location:     gcs.Metadata.Labels.Location,
		CreationTime: createTime.Local(),
	}
}

type Function struct {
	Name        string
	Location    string
	Environment string
	Runtime     string
	State       string
	UpdateTime  time.Time
}

func (f Function) Title() string {
	return f.Name
}

func (f Function) Subtitle() string {
	var icon string
	switch f.State {
	case "ACTIVE":
		icon = "🟢"
	default:
		icon = "❓"
	}

	var version string
	if f.Environment == "GEN_1" {
		version = "Gen 1 aka functions"
	} else {
		version = "Gen 2 aka services"
	}

	return fmt.Sprintf("%s %s | %s | Last Updated: %s",
		icon, version, f.Runtime, util.FormatTime(f.UpdateTime))
}

func (f Function) URL(config *gc.Config) string {
	if f.Environment == "GEN_1" {
		return fmt.Sprintf(
			"https://console.cloud.google.com/functions/details/%s/%s?env=gen1&project=%s",
			f.Location,
			f.Name,
			config.Project,
		)
	}
	// gen2 url
	return fmt.Sprintf(
		"https://console.cloud.google.com/run/detail/%s/%s/metrics?project=%s",
		f.Location,
		f.Name,
		config.Project,
	)
}

func FromGCloudCloudRunFunction(gcf *gc.CloudRunFunction) Function {
	var name string
	var location string

	parts := strings.Split(gcf.Name, "/")
	for i, part := range parts {
		switch part {
		case "locations":
			if i+1 < len(parts) {
				location = parts[i+1]
			}
		case "functions":
			if i+1 < len(parts) {
				name = parts[i+1]
			}
		}
	}

	var updateTime time.Time
	updateTime, _ = time.Parse(time.RFC3339, gcf.UpdateTime)

	return Function{
		Name:        name,
		Location:    location,
		Environment: gcf.Environment,
		Runtime:     gcf.BuildConfig.Runtime,
		State:       gcf.State,
		UpdateTime:  updateTime.Local(),
	}
}
