package netconnectivity

import (
	"fmt"
	"strings"

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
