package gcloud

import "strings"

type Region struct {
	Name     string
	Location string
	Zones    []string
}

// Static list of GCP regions with metadata.
// For some weird reason the `gcloud compute regions list --format=json` does not return location name :sad
// So we run a script in ./tools/update-regions.js to get a new list of regions
var regions = []Region{
	{
		Name:     "africa-south1",
		Location: "Johannesburg, South Africa",
		Zones:    []string{"africa-south1-a", "africa-south1-b", "africa-south1-c"},
	},
	{
		Name:     "asia-east1",
		Location: "Changhua County, Taiwan, APAC",
		Zones:    []string{"asia-east1-a", "asia-east1-b", "asia-east1-c"},
	},
	{
		Name:     "asia-east2",
		Location: "Hong Kong, APAC",
		Zones:    []string{"asia-east2-a", "asia-east2-b", "asia-east2-c"},
	},
	{
		Name:     "asia-northeast1",
		Location: "Tokyo, Japan, APAC",
		Zones:    []string{"asia-northeast1-a", "asia-northeast1-b", "asia-northeast1-c"},
	},
	{
		Name:     "asia-northeast2",
		Location: "Osaka, Japan, APAC",
		Zones:    []string{"asia-northeast2-a", "asia-northeast2-b", "asia-northeast2-c"},
	},
	{
		Name:     "asia-northeast3",
		Location: "Seoul, South Korea, APAC",
		Zones:    []string{"asia-northeast3-a", "asia-northeast3-b", "asia-northeast3-c"},
	},
	{
		Name:     "asia-south1",
		Location: "Mumbai, India, APAC",
		Zones:    []string{"asia-south1-a", "asia-south1-b", "asia-south1-c"},
	},
	{
		Name:     "asia-south2",
		Location: "Delhi, India, APAC",
		Zones:    []string{"asia-south2-a", "asia-south2-b", "asia-south2-c"},
	},
	{
		Name:     "asia-southeast1",
		Location: "Jurong West, Singapore, APAC",
		Zones:    []string{"asia-southeast1-a", "asia-southeast1-b", "asia-southeast1-c"},
	},
	{
		Name:     "asia-southeast2",
		Location: "Jakarta, Indonesia, APAC",
		Zones:    []string{"asia-southeast2-a", "asia-southeast2-b", "asia-southeast2-c"},
	},
	{
		Name:     "australia-southeast1",
		Location: "Sydney, Australia, APAC",
		Zones:    []string{"australia-southeast1-a", "australia-southeast1-b", "australia-southeast1-c"},
	},
	{
		Name:     "australia-southeast2",
		Location: "Melbourne, Australia, APAC",
		Zones:    []string{"australia-southeast2-a", "australia-southeast2-b", "australia-southeast2-c"},
	},
	{
		Name:     "europe-central2",
		Location: "Warsaw, Poland, Europe",
		Zones:    []string{"europe-central2-a", "europe-central2-b", "europe-central2-c"},
	},
	{
		Name:     "europe-north1",
		Location: "Hamina, Finland, Europe",
		Zones:    []string{"europe-north1-a", "europe-north1-b", "europe-north1-c"},
	},
	{
		Name:     "europe-north2",
		Location: "Stockholm, Sweden, Europe",
		Zones:    []string{"europe-north2-a", "europe-north2-b", "europe-north2-c"},
	},
	{
		Name:     "europe-southwest1",
		Location: "Madrid, Spain, Europe",
		Zones:    []string{"europe-southwest1-a", "europe-southwest1-b", "europe-southwest1-c"},
	},
	{
		Name:     "europe-west1",
		Location: "St. Ghislain, Belgium, Europe",
		Zones:    []string{"europe-west1-b", "europe-west1-c", "europe-west1-d"},
	},
	{
		Name:     "europe-west10",
		Location: "Berlin, Germany, Europe",
		Zones:    []string{"europe-west10-a", "europe-west10-b", "europe-west10-c"},
	},
	{
		Name:     "europe-west12",
		Location: "Turin, Italy, Europe",
		Zones:    []string{"europe-west12-a", "europe-west12-b", "europe-west12-c"},
	},
	{
		Name:     "europe-west2",
		Location: "London, England, Europe",
		Zones:    []string{"europe-west2-a", "europe-west2-b", "europe-west2-c"},
	},
	{
		Name:     "europe-west3",
		Location: "Frankfurt, Germany, Europe",
		Zones:    []string{"europe-west3-a", "europe-west3-b", "europe-west3-c"},
	},
	{
		Name:     "europe-west4",
		Location: "Eemshaven, Netherlands, Europe",
		Zones:    []string{"europe-west4-a", "europe-west4-b", "europe-west4-c"},
	},
	{
		Name:     "europe-west6",
		Location: "Zurich, Switzerland, Europe",
		Zones:    []string{"europe-west6-a", "europe-west6-b", "europe-west6-c"},
	},
	{
		Name:     "europe-west8",
		Location: "Milan, Italy, Europe",
		Zones:    []string{"europe-west8-a", "europe-west8-b", "europe-west8-c"},
	},
	{
		Name:     "europe-west9",
		Location: "Paris, France, Europe",
		Zones:    []string{"europe-west9-a", "europe-west9-b", "europe-west9-c"},
	},
	{
		Name:     "me-central1",
		Location: "Doha, Qatar, Middle East",
		Zones:    []string{"me-central1-a", "me-central1-b", "me-central1-c"},
	},
	{
		Name:     "me-central2",
		Location: "Dammam, Saudi Arabia, Middle East",
		Zones:    []string{"me-central2-a", "me-central2-b", "me-central2-c"},
	},
	{
		Name:     "me-west1",
		Location: "Tel Aviv, Israel, Middle East",
		Zones:    []string{"me-west1-a", "me-west1-b", "me-west1-c"},
	},
	{
		Name:     "northamerica-northeast1",
		Location: "Montréal, Québec, North America",
		Zones:    []string{"northamerica-northeast1-a", "northamerica-northeast1-b", "northamerica-northeast1-c"},
	},
	{
		Name:     "northamerica-northeast2",
		Location: "Toronto, Ontario, North America",
		Zones:    []string{"northamerica-northeast2-a", "northamerica-northeast2-b", "northamerica-northeast2-c"},
	},
	{
		Name:     "northamerica-south1",
		Location: "Queretaro, Mexico, North America",
		Zones:    []string{"northamerica-south1-a", "northamerica-south1-b", "northamerica-south1-c"},
	},
	{
		Name:     "southamerica-east1",
		Location: "Osasco, São Paulo, Brazil, South America",
		Zones:    []string{"southamerica-east1-a", "southamerica-east1-b", "southamerica-east1-c"},
	},
	{
		Name:     "southamerica-west1",
		Location: "Santiago, Chile, South America",
		Zones:    []string{"southamerica-west1-a", "southamerica-west1-b", "southamerica-west1-c"},
	},
	{
		Name:     "us-central1",
		Location: "Council Bluffs, Iowa, North America",
		Zones:    []string{"us-central1-a", "us-central1-b", "us-central1-c", "us-central1-f"},
	},
	{
		Name:     "us-east1",
		Location: "Moncks Corner, South Carolina, North America",
		Zones:    []string{"us-east1-b", "us-east1-c", "us-east1-d"},
	},
	{
		Name:     "us-east4",
		Location: "Ashburn, Virginia, North America",
		Zones:    []string{"us-east4-a", "us-east4-b", "us-east4-c"},
	},
	{
		Name:     "us-east5",
		Location: "Columbus, Ohio, North America",
		Zones:    []string{"us-east5-a", "us-east5-b", "us-east5-c"},
	},
	{
		Name:     "us-south1",
		Location: "Dallas, Texas, North America",
		Zones:    []string{"us-south1-a", "us-south1-b", "us-south1-c"},
	},
	{
		Name:     "us-west1",
		Location: "The Dalles, Oregon, North America",
		Zones:    []string{"us-west1-a", "us-west1-b", "us-west1-c"},
	},
	{
		Name:     "us-west2",
		Location: "Los Angeles, California, North America",
		Zones:    []string{"us-west2-a", "us-west2-b", "us-west2-c"},
	},
	{
		Name:     "us-west3",
		Location: "Salt Lake City, Utah, North America",
		Zones:    []string{"us-west3-a", "us-west3-b", "us-west3-c"},
	},
	{
		Name:     "us-west4",
		Location: "Las Vegas, Nevada, North America",
		Zones:    []string{"us-west4-a", "us-west4-b", "us-west4-c"},
	},
}

func GetAllRegionNames() []string {
	out := make([]string, 0, len(regions))
	for _, r := range regions {
		out = append(out, r.Name)
	}
	return out
}

func GetAllRegions() []Region {
	return regions
}

func GetRegionByName(name string) *Region {
	for _, r := range regions {
		if strings.EqualFold(r.Name, name) {
			return &r
		}
	}
	return nil
}
