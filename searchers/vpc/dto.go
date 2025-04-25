package vpc

import (
	"fmt"
	"time"

	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Network struct {
	Id           string
	Name         string
	Description  string
	CreationTime time.Time
	RoutingMode  string
}

func (n Network) Title() string {
	return n.Name
}

func (n Network) Subtitle() string {
	return fmt.Sprintf(" %s | Created: %s", n.RoutingMode, n.CreationTime.Format("Jan 2, 2006 15:04 MST"))
}

func (n Network) URL(config *gc.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/networking/networks/details/%s?project=%s", n.Name, config.Project)
}

func NetworkFromGCloud(network *gc.VPCNetwork) Network {
	createdAt, err := time.Parse(time.RFC3339, network.CreationTime)
	if err != nil {
		createdAt = time.Time{}
	}

	return Network{
		Id:           network.Id,
		Name:         network.Name,
		Description:  network.Description,
		CreationTime: createdAt.Local(),
		RoutingMode:  network.XCloudBgpRoutingMode,
	}
}
