package vpc

import (
	"fmt"
	"strings"
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

func FromGCloudNetwork(network *gc.VPCNetwork) Network {
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

type Route struct {
	Id           string
	Name         string
	Network      string
	Priority     int
	CreationTime time.Time
}

func (r Route) Title() string {
	return r.Name
}

func (r Route) Subtitle() string {
	return fmt.Sprintf(" %s | Priority: %d | Created: %s", r.Network, r.Priority, r.CreationTime.Format("Jan 2, 2006 15:04 MST"))
}

func (r Route) URL(config *gc.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/networking/routes/details/%s?project=%s", r.Name, config.Project)
}

func FromGCloudRoute(route *gc.VPCRoute) Route {
	createdAt, err := time.Parse(time.RFC3339, route.CreationTime)
	if err != nil {
		createdAt = time.Time{}
	}

	var network string
	if route.Network != "" {
		parts := strings.Split(route.Network, "/")
		for i, part := range parts {
			if part == "networks" && i+1 < len(parts) {
				network = parts[i+1]
				break
			}
		}
	}

	return Route{
		Id:           route.Id,
		Name:         route.Name,
		Network:      network,
		Priority:     route.Priority,
		CreationTime: createdAt.Local(),
	}
}
