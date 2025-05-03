package gcloud

import (
	"embed"
	"strings"

	"gopkg.in/yaml.v2"
)

type Region struct {
	Name     string
	Location string
	Zones    []string
}

// Static list of GCP regions
// For some weird reason the `gcloud compute regions list --format=json` does not return location name :sad
// In order to populate new list check out tools/update-regions.js and update the regions.yml file
var regions []Region

// InitRegions initializes the regions list from the embedded file system.
// This function should be called only once, typically at the start of the program.
func InitRegions(fs embed.FS) {
	if len(regions) > 0 {
		return
	}

	data, err := fs.ReadFile("regions.yml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &regions); err != nil {
		// no point in continuing if we can't parse the regions
		panic(err)
	}
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
