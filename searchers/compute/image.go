package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type ImageSearcher struct{}

func (s *ImageSearcher) Search(wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result) error {
	builder := resource.NewBuilder(
		"compute_images",
		wf,
		cfg,
		q,
		gc.ListComputeImages,
		func(wf *aw.Workflow, gci gc.ComputeImage) {
			ci := FromGCloudComputeImage(&gci)
			resource.NewItem(wf, cfg, ci, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
