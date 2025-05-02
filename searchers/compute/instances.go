package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type InstanceSearcher struct{}

func (s *InstanceSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"compute_instances",
		wf,
		cfg,
		q,
		gc.ListComputeInstances,
		func(wf *aw.Workflow, gci gc.ComputeInstance) {
			ci := FromGCloudComputeInstance(&gci)
			resource.NewItem(wf, cfg, ci, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
