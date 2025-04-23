package services

import (
	"bytes"
	"embed"
	"text/template"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"gopkg.in/yaml.v2"
)

// Service defines the structure of each GCP service
type Service struct {
	ID          string    `yaml:"id"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	URL         string    `yaml:"url"`
	SubServices []Service `yaml:"sub_services"`
	LogoPath    string    `yaml:"logo_path"`
}

func (s *Service) Icon() *aw.Icon {
	return &aw.Icon{
		Value: "/Users/dinesh.chikkanna/personal/alfred-gcp-workflow/" + s.LogoPath,
	}
}

func (s *Service) Url(config *gcloud.Config) (string, error) {
	t, err := template.New("url").Parse(s.URL)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, config); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *Service) Subtitle() string {
	if len(s.SubServices) > 0 {
		return "📁 " + s.Description
	}

	return s.Description
}

// Load reads the services from an embedded YAML file
func Load(file embed.FS) ([]Service, error) {
	data, err := file.ReadFile("services.yml")
	if err != nil {
		return nil, err
	}

	var services []Service
	if err := yaml.Unmarshal(data, &services); err != nil {
		return nil, err
	}

	return services, nil
}
