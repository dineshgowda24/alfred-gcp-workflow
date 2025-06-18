package monitoring

import (
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Dashboard struct {
	DisplayName string
	Name        string
}

func (d Dashboard) Title() string {
	return d.DisplayName
}

func (d Dashboard) Subtitle() string {
	return d.ID()
}

func (d Dashboard) URL(config *gcloud.Config) string {
	id := d.ID()
	return "https://console.cloud.google.com/monitoring/dashboards/builder/" + id + "?project=" + config.Project

}

func (d Dashboard) ID() string {
	parts := strings.Split(d.Name, "/")
	return parts[len(parts)-1]
}

func FromGCloudMonitoringDashboard(board *gcloud.Dashboard) Dashboard {
	return Dashboard{
		DisplayName: board.DisplayName,
		Name:        board.Name,
	}
}
