package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type MachineImageSearcher struct{}

func (s *MachineImageSearcher) Search(
	wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result,
) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"compute_machine_images",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.ComputeMachineImage) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *MachineImageSearcher) fetch(config *gc.Config) ([]gc.ComputeMachineImage, error) {
	return gc.ListComputeMachineImages(config)
}

func (s *MachineImageSearcher) render(
	wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.ComputeMachineImage,
) {
	image := FromGCloudComputeMachineImage(&entity)
	wf.NewItem(image.Title()).
		Subtitle(image.Subtitle()).
		Arg(image.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
