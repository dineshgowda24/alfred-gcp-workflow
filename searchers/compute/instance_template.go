package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type InstanceTmplSearcher struct{}

func (s *InstanceTmplSearcher) Search(
	wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result,
) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"compute_instance_templates",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.ComputeInstanceTemplate) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *InstanceTmplSearcher) fetch(config *gc.Config) ([]gc.ComputeInstanceTemplate, error) {
	return gc.ListComputeInstanceTemplates(config)
}

func (s *InstanceTmplSearcher) render(
	wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.ComputeInstanceTemplate,
) {
	tmpl := FromGCloudComputeInstanceTemplate(&entity)
	wf.NewItem(tmpl.Title()).
		Subtitle(tmpl.Subtitle()).
		Arg(tmpl.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
