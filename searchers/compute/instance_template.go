package compute

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type InstanceTemplateSearcher struct{}

func (s *InstanceTemplateSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, filter string) error {
	gtmpls, err := gcloud.ListComputeInstanceTemplates(config)
	if err != nil {
		return err
	}

	for _, gtmpl := range gtmpls {
		tmpl := FromGCloudComputeInstanceTemplate(&gtmpl)
		wf.NewItem(tmpl.Title()).
			Subtitle(tmpl.Subtitle()).
			Arg(tmpl.URL(config)).
			Icon(svc.Icon()).
			Valid(true)
	}

	if filter != "" {
		wf.Filter(filter)
	}

	return nil
}
