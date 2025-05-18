package gcloud

import "strings"

type CloudRunService struct {
	Metadata struct {
		Name   string `json:"name"`
		Labels struct {
			Location string `json:"cloud.googleapis.com/location"`
		} `json:"labels"`
		CreationTimestamp string `json:"creationTimestamp"`
	}
}

func (rs CloudRunService) IsRegionSupported(config *Config) bool {
	if config == nil || config.GetRegionName() == "" {
		return true
	}

	type Region struct {
		LocationID string `json:"locationId"`
	}

	regions, err := runGCloudCmd[[]Region](config, "run", "regions", "list", "--format=json(locationId)")
	if err != nil {
		return false
	}

	for _, r := range regions {
		if strings.EqualFold(r.LocationID, config.GetRegionName()) {
			return true
		}
	}

	return false
}

func ListCloudRunServices(config *Config) ([]CloudRunService, error) {
	args := []string{"services", "list", "--format=json(metadata.name,metadata.labels,metadata.creationTimestamp)"}
	if config != nil && config.GetRegionName() != "" {
		args = append(args, "--region="+config.GetRegionName())
	}

	return runGCloudCmd[[]CloudRunService](config, "run", args...)
}

// https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services
// Since we are using --v2 flag while listing, we can refer the same services reference
type CloudRunFunction struct {
	Name        string `json:"name"`
	State       string `json:"state"`
	Environment string `json:"environment"`
	UpdateTime  string `json:"updateTime"`
	BuildConfig struct {
		Runtime string `json:"runtime"`
	} `json:"buildConfig"`
}

func (rf CloudRunFunction) IsRegionSupported(config *Config) bool {
	if config == nil || config.GetRegionName() == "" {
		return true
	}

	type Region struct {
		LocationID string `json:"locationId"`
	}

	regions, err := runGCloudCmd[[]Region](config, "functions", "regions", "list", "--format=json(locationId)")
	if err != nil {
		return false
	}

	for _, r := range regions {
		if strings.EqualFold(r.LocationID, config.GetRegionName()) {
			return true
		}
	}

	return false
}

// ListCloudRunFunctions retrieves only Cloud Functions (Gen 1).
// By default, `gcloud functions list --v2` includes both Gen 1 and Gen 2,
// so we apply a filter on `environment=GEN_1` to exclude Gen 2 services.
// The --v2 flag ensures a consistent data format across environments.
func ListCloudRunFunctions(config *Config) ([]CloudRunFunction, error) {
	args := []string{
		"list", "--v2",
		"--format=json(name,state,environment,updateTime,buildConfig.runtime)", "--filter=environment=GEN_1",
	}

	if config != nil && config.GetRegionName() != "" {
		args = append(args, "--regions="+config.GetRegionName())
	}
	return runGCloudCmd[[]CloudRunFunction](config, "functions", args...)
}
