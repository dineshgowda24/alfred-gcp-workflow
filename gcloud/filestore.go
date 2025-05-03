package gcloud

import "strings"

// https://cloud.google.com/filestore/docs/reference/rest/v1/projects.locations.instances
type FilestoreInstance struct {
	Name       string `json:"name"`
	State      string `json:"state"`
	Tier       string `json:"tier"`
	FileShares []struct {
		CapacityGb string `json:"capacityGb"`
		Name       string `json:"name"`
	} `json:"fileShares"`
	CreatedAt string `json:"createTime"`
}

func (fi FilestoreInstance) IsRegionSupported(config *Config) bool {
	if config == nil || config.GetRegionName() == "" {
		return true
	}

	type Alias struct {
		LocationId string `json:"locationId"`
	}

	regions, err := runGCloudCmd[[]Alias](config, "filestore", "regions", "list", "--format=json(locationId)")
	if err != nil {
		return false
	}

	for _, r := range regions {
		if strings.EqualFold(r.LocationId, config.GetRegionName()) {
			return true
		}
	}

	return false
}

func ListFilestoreInstances(config *Config) ([]FilestoreInstance, error) {
	args := []string{"instances", "list", "--format=json(name,state,tier,fileShares,createTime)"}
	if config != nil && config.GetRegionName() != "" {
		args = append(args, "--region="+config.GetRegionName())
	}

	return runGCloudCmd[[]FilestoreInstance](config, "filestore", args...)
}
