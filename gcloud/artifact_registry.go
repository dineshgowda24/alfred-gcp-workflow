package gcloud

// https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories#Repository
type ArtifactRepository struct {
	Name       string `json:"name"`
	Format     string `json:"format"`
	UpdateTime string `json:"updateTime"`
}

func ListArtifactRepositories(config *Config) ([]ArtifactRepository, error) {
	return runGCloudCmd[[]ArtifactRepository](config,
		"artifacts", "repositories", "list",
		"--format=json(name,format,updateTime)",
	)
}
