package netconnectivity

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type VPNTunnel struct {
	Name    string
	Status  string
	Region  string
	Gateway string
}

func (t *VPNTunnel) Title() string {
	return t.Name
}

func (t *VPNTunnel) Subtitle() string {
	var icon string
	switch t.Status {
	case "ESTABLISHED", "FIRST_HANDSHAKE":
		icon = "🟢"
	case "PROVISIONING", "WAITING_FOR_FULL_CONFIG", "ALLOCATING_RESOURCES":
		icon = "🕒"
	case "NETWORK_ERROR", "AUTHORIZATION_ERROR", "NEGOTIATION_FAILURE", "DEPROVISIONING",
		"FAILED", "NO_INCOMING_PACKETS", "REJECTED", "STOPPED", "PEER_IDENTITY_MISMATCH",
		"TS_NARROWING_NOT_ALLOWED":
		icon = "❌"
	default:
		icon = "❓"
	}

	return fmt.Sprintf("%s Gateway: %s", icon, t.Gateway)
}

func (t *VPNTunnel) URL(config *gcloud.Config) string {
	return fmt.Sprintf("https://console.cloud.google.com/hybrid/vpn/tunnels/details/%s/%s?project=%s",
		t.Region, t.Name, config.Project)
}

func VPNTunnelFromGCloud(tunnel *gcloud.VPNTunnel) VPNTunnel {
	var region string
	words := strings.Split(tunnel.Region, "/")
	for i, word := range words {
		if word == "regions" && i+1 < len(words) {
			region = words[i+1]
			break
		}
	}

	var gateway string
	words = strings.Split(tunnel.VPNGateway, "/")
	for i, word := range words {
		if word == "vpnGateways" && i+1 < len(words) {
			gateway = words[i+1]
			break
		}
	}

	return VPNTunnel{
		Name:    tunnel.Name,
		Status:  tunnel.Status,
		Region:  region,
		Gateway: gateway,
	}
}

type VPNGateway struct {
	Name             string
	GatewayIPVersion string
	Network          string
	Region           string
	CreationTime     time.Time
}

func (g *VPNGateway) Title() string {
	return g.Name
}

func (g *VPNGateway) Subtitle() string {
	return fmt.Sprintf("IP Version: %s | Network: %s | Created: %s",
		g.GatewayIPVersion, g.Network, g.CreationTime.Format("Jan 2, 2006 15:04 MST"))
}

func (g *VPNGateway) URL(config *gcloud.Config) string {
	return fmt.Sprintf(
		"https://console.cloud.google.com/hybrid/vpn/gateways/details/%s/%s?project=%s",
		g.Region, g.Name, config.Project)
}

func VPNGatewayFromGCloud(gateway *gcloud.VPNGateway) VPNGateway {
	var region string
	words := strings.Split(gateway.Region, "/")
	for i, word := range words {
		if word == "regions" && i+1 < len(words) {
			region = words[i+1]
			break
		}
	}

	var network string
	words = strings.Split(gateway.Network, "/")
	for i, word := range words {
		if word == "networks" && i+1 < len(words) {
			network = words[i+1]
			break
		}
	}

	var creationTime time.Time
	creationTime, err := time.Parse(time.RFC3339, gateway.CreationTimestamp)
	if err != nil {
		log.Println("LOG: Error parsing creation timestamp:", err)
	}

	return VPNGateway{
		Name:             gateway.Name,
		GatewayIPVersion: gateway.GatewayIPVersion,
		Network:          network,
		Region:           region,
		CreationTime:     creationTime.Local(),
	}
}
