package gcloud

import "fmt"

type VPNTunnel struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Region     string `json:"region"`
	VPNGateway string `json:"vpnGateway"`
}

func (t VPNTunnel) IsRegionSupported(config *Config) bool {
	return isComputeRegionSupported(config)
}

// --regions flag for some reason doesn't work and always returns all regions even if you specify one
// so we have use filter instead
func ListVPNTunnels(config *Config) ([]VPNTunnel, error) {
	args := []string{"vpn-tunnels", "list", "--format=json(name,status,region,vpnGateway)"}
	if config != nil && config.GetRegionName() != "" {
		args = append(args, fmt.Sprintf("--filter=region:(%s)", config.GetRegionName()))
	}
	return runGCloudCmd[[]VPNTunnel](config, "compute", args...)
}

type VPNGateway struct {
	Name              string `json:"name"`
	GatewayIPVersion  string `json:"gatewayIpVersion"`
	Network           string `json:"network"`
	Region            string `json:"region"`
	CreationTimestamp string `json:"creationTimestamp"`
}

func (g VPNGateway) IsRegionSupported(config *Config) bool {
	return isComputeRegionSupported(config)
}

// --regions flag for some reason doesn't work and always returns all regions even if you specify one
// so we have use filter instead
func ListVPNGateways(config *Config) ([]VPNGateway, error) {
	args := []string{"vpn-gateways", "list", "--format=json(name,gatewayIpVersion,network,region,creationTimestamp)"}
	if config != nil && config.GetRegionName() != "" {
		args = append(args, fmt.Sprintf("--filter=region:(%s)", config.GetRegionName()))
	}
	return runGCloudCmd[[]VPNGateway](config, "compute", args...)
}

type CloudRouter struct {
	Name              string `json:"name"`
	Network           string `json:"network"`
	Region            string `json:"region"`
	CreationTimestamp string `json:"creationTimestamp"`
}

func (r CloudRouter) IsRegionSupported(config *Config) bool {
	return isComputeRegionSupported(config)
}

// --regions flag for some reason doesn't work and always returns all regions even if you specify one
// so we have use filter instead
func ListCloudRouters(config *Config) ([]CloudRouter, error) {
	args := []string{"routers", "list", "--format=json(name,network,region,creationTimestamp)"}
	if config != nil && config.GetRegionName() != "" {
		args = append(args, fmt.Sprintf("--filter=region:(%s)", config.GetRegionName()))
	}
	return runGCloudCmd[[]CloudRouter](config, "compute", args...)
}
