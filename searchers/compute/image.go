package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type ImageSearcher struct{}

func (s *ImageSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"compute_images",
		wf,
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gc.ComputeImage) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *ImageSearcher) fetch(config *gc.Config) ([]gc.ComputeImage, error) {
	return gc.ListComputeImages(config)
}

func (s *ImageSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.ComputeImage) {
	image := FromGCloudComputeImage(&entity)
	wf.NewItem(image.Title()).
		Subtitle(image.Subtitle()).
		Arg(image.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
