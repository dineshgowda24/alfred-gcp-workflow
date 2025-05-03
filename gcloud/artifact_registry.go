package gcloud

import "strings"

// https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories#Repository
type ArtifactRepository struct {
	Name       string `json:"name"`
	Format     string `json:"format"`
	UpdateTime string `json:"updateTime"`
}

func (r ArtifactRepository) IsRegionSupported(config *Config) bool {
	if config == nil || config.GetRegionName() == "" {
		return true
	}

	type Location struct {
		Name string `json:"name"`
	}

	locations, err := runGCloudCmd[[]Location](config, "artifacts", "locations", "list", "--format=json(name)")
	if err != nil {
		return false
	}

	for _, loc := range locations {
		if strings.EqualFold(loc.Name, config.GetRegionName()) {
			return true
		}
	}

	return false
}

func ListArtifactRepositories(config *Config) ([]ArtifactRepository, error) {
	args := []string{"artifacts", "repositories", "list", "--format=json(name,format,updateTime)"}

	if config != nil && config.GetRegionName() != "" {
		args = append(args, "--location="+config.GetRegionName())
	}
	return runGCloudCmd[[]ArtifactRepository](config, "artifacts", args...)
}
