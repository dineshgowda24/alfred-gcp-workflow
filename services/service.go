package services

import (
	"bytes"
	"embed"
	"path/filepath"
	"text/template"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/log"
	"gopkg.in/yaml.v2"
)

type Exister interface {
	Exists(parent, child *Service) bool
}

// Service defines the structure of each GCP service
type Service struct {
	ID          string    `yaml:"id"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	URL         string    `yaml:"url"`
	SubServices []Service `yaml:"sub_services"`
	LogoPath    string    `yaml:"logo_path"`

	// Parent points to the parent Service in the hierarchy.
	// It is set after unmarshalling to allow navigating upward in the service tree.
	Parent *Service `yaml:"-"`
}

func (s *Service) Title() string {
	if s.IsChild() {
		return s.Parent.ID + " " + s.ID
	}
	return s.ID
}

func (s *Service) Autocomplete() string {
	if s.IsChild() {
		return s.Parent.ID + " " + s.ID
	}

	return s.ID
}

func (s *Service) Icon() *aw.Icon {
	return &aw.Icon{
		Value: filepath.Clean(s.LogoPath),
	}
}

func (s *Service) Url(config *gcloud.Config) (string, error) {
	t, err := template.New("url").Parse(s.URL)
	if err != nil {
		log.Errorf("error parsing template for service %s: %v", s.Name, err)
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, config); err != nil {
		log.Errorf("error executing template for service %s: %v", s.Name, err)
		return "", err
	}
	return buf.String(), nil
}

func (s *Service) Subtitle(exister Exister) string {
	if s.IsParent() {
		suffix := s.Name + " - " + s.Description
		if len(s.SubServices) > 0 {
			return "🗂️ " + suffix
		}
		return suffix
	} else if s.IsChild() {
		suffix := s.Parent.Name + " - " + s.Name
		if exister != nil && exister.Exists(s.Parent, s) {
			return "🔍⚡️ " + suffix
		}
		return suffix
	}
	return s.Description
}

func (s *Service) Match() string {
	return s.ID + " " + s.Name
}

func (s *Service) IsParent() bool {
	return s.Parent == nil
}

func (s *Service) IsChild() bool {
	return s.Parent != nil
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

	linkParentChild(services)
	return services, nil
}

func linkParentChild(services []Service) {
	for i := range services {
		parent := &services[i]
		for j := range parent.SubServices {
			child := &parent.SubServices[j]
			child.Parent = parent
			linkParentChild(child.SubServices)
		}
	}
}
