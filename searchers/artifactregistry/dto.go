package artifactregistry

import (
	"fmt"
	"strings"
	"time"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Repository struct {
	Name       string
	Format     string
	Location   string
	UpdateTime time.Time
}

func (r *Repository) Title() string {
	return r.Name
}

func (r *Repository) Subtitle() string {
	return fmt.Sprintf("Format: %s  | Last Updated: %s", r.Format, r.UpdateTime.Format("Jan 2, 2006 15:04 MST"))
}

func (r *Repository) URL(config *gcloud.Config) string {
	format := strings.ToLower(r.Format)
	return fmt.Sprintf(
		"https://console.cloud.google.com/artifacts/%s/%s/%s/%s?project=%s",
		format, config.Project, r.Location, r.Name, config.Project)
}

func RepositoryFromGCloud(repo *gcloud.ArtifactRepository) Repository {
	var updatedTime time.Time
	updatedTime, err := time.Parse(time.RFC3339, repo.UpdateTime)
	if err != nil {
		updatedTime = time.Time{}
	}

	var location string
	var name string
	words := strings.Split(repo.Name, "/")
	for i, word := range words {
		switch strings.ToLower(word) {
		case "locations":
			if i+1 < len(words) {
				location = words[i+1]
			}
		case "repositories":
			if i+1 < len(words) {
				name = words[i+1]
			}
		}
	}

	return Repository{
		Name:       name,
		Format:     repo.Format,
		Location:   location,
		UpdateTime: updatedTime.Local(),
	}
}
