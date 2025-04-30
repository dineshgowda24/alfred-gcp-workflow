package netservices

import (
	"fmt"
	"time"

	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type DNSZone struct {
	Id           string
	Name         string
	DnsName      string
	CreationTime time.Time
	Visibility   string
}

func (d DNSZone) Title() string {
	return d.Name
}

func (d DNSZone) Subtitle() string {
	var icon string
	switch d.Visibility {
	case "public":
		icon = "🛰️"
	case "private":
		icon = "🛡️"
	default:
		icon = "❓"
	}

	return icon + " " + d.DnsName + " | " + d.CreationTime.Format("Jan 2, 2006 15:04 MST")
}

func (d DNSZone) URL(config *gc.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/net-services/dns/zones/%s/details?project=%s", d.Name, config.Project)
}

func FromGCloudDNSZone(zone *gc.DNSZone) DNSZone {
	createdAt, err := time.Parse(time.RFC3339, zone.CreationTime)
	if err != nil {
		createdAt = time.Time{}
	}

	return DNSZone{
		Id:           zone.Id,
		Name:         zone.Name,
		DnsName:      zone.DnsName,
		CreationTime: createdAt,
		Visibility:   zone.Visibility,
	}
}
