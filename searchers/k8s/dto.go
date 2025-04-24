package k8s

import (
	"fmt"
	"time"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Cluster struct {
	Name          string
	Status        string
	Location      string
	MasterVersion string
	NodeCount     int
	CreatedAt     time.Time
}

func (c Cluster) Title() string {
	return c.Name
}

func (c Cluster) Subtitle() string {
	var icon string
	switch c.Status {
	case "RUNNING":
		icon = "🟢"
	case "PROVISIONING", "RECONCILING":
		icon = "🕒"
	case "STOPPING", "DEGRADED":
		icon = "❌"
	case "ERROR":
		icon = "❗"
	default:
		icon = "❓"
	}

	return fmt.Sprintf(
		"%s %s | Nodes: %d | %s",
		icon, c.MasterVersion, c.NodeCount, c.CreatedAt.Local().Format("Jan 2, 2006 15:04 MST"),
	)
}

func (c Cluster) URL(config *gcloud.Config) string {
	return fmt.Sprintf(
		"https://console.cloud.google.com/kubernetes/clusters/details/%s/%s?project=%s",
		c.Location, c.Name, config.Project,
	)
}

func ClusterFromGCloud(cluster *gcloud.K8sCluster) Cluster {
	createdAt, err := time.Parse("2006-01-02T15:04:05-07:00", cluster.CreatedAt)
	if err != nil {
		createdAt = time.Time{}
	}

	return Cluster{
		Name:          cluster.Name,
		Status:        cluster.Status,
		Location:      cluster.Location,
		MasterVersion: cluster.CurrentMasterVersion,
		NodeCount:     cluster.CurrentNodeCount,
		CreatedAt:     createdAt,
	}
}
