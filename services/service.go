package services

import (
	"embed"

	aw "github.com/deanishe/awgo"
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

func LoadServices(file embed.FS) ([]Service, error) {
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
