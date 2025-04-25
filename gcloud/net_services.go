package gcloud

type DNSZone struct {
	CreationTime string `json:"creationTime"`
	DnsName      string `json:"dnsName"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Visibility   string `json:"visibility"`
}

func ListDNSZones(config *Config) ([]DNSZone, error) {
	return runGCloudCmd[[]DNSZone](
		config,
		"dns", "managed-zones", "list",
		"--format=json(creationTime,description,dnsName,id,name,visibility)",
	)
}
