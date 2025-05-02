package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type MachineImageSearcher struct{}

func (s *MachineImageSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"compute_machine_images",
		wf,
		cfg,
		q,
		gc.ListComputeMachineImages,
		func(wf *aw.Workflow, gcmi gc.ComputeMachineImage) {
			cmi := FromGCloudComputeMachineImage(&gcmi)
			resource.NewItem(wf, cfg, cmi, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
