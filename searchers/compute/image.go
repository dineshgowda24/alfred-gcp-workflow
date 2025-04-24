package compute

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type ImageSearcher struct{}

func (s *ImageSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, args workflow.SearchArgs) error {
	return workflow.LoadFromCache(
		wf,
		config.CacheKey("compute_images"),
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gcloud.ComputeImage) {
			s.render(wf, svc, config, entity)
		},
	)
}

func (s *ImageSearcher) fetch(config *gcloud.Config) ([]gcloud.ComputeImage, error) {
	return gcloud.ListComputeImages(config)
}

func (s *ImageSearcher) render(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, entity gcloud.ComputeImage) {
	image := FromGCloudComputeImage(&entity)
	wf.NewItem(image.Title()).
		Subtitle(image.Subtitle()).
		Arg(image.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
