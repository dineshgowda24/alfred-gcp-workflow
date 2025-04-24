package compute

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type InstanceTemplateSearcher struct{}

func (s *InstanceTemplateSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, args workflow.SearchArgs) error {
	return workflow.LoadFromCache(
		wf,
		config.CacheKey("compute_instance_templates"),
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gcloud.ComputeInstanceTemplate) {
			s.render(wf, svc, config, entity)
		},
	)
}

func (s *InstanceTemplateSearcher) fetch(config *gcloud.Config) ([]gcloud.ComputeInstanceTemplate, error) {
	return gcloud.ListComputeInstanceTemplates(config)
}

func (s *InstanceTemplateSearcher) render(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, entity gcloud.ComputeInstanceTemplate) {
	tmpl := FromGCloudComputeInstanceTemplate(&entity)
	wf.NewItem(tmpl.Title()).
		Subtitle(tmpl.Subtitle()).
		Arg(tmpl.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
