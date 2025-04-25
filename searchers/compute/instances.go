package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type InstanceSearcher struct{}

func (s *InstanceSearcher) Search(
	wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result,
) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"compute_instances",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.ComputeInstance) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *InstanceSearcher) fetch(config *gc.Config) ([]gc.ComputeInstance, error) {
	return gc.ListComputeInstances(config)
}

func (s *InstanceSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.ComputeInstance) {
	inst := FromGCloudComputeInstance(&entity)
	wf.NewItem(inst.Title()).
		Subtitle(inst.Subtitle()).
		Arg(inst.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
