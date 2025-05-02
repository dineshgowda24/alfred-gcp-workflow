package cloudrun

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type FunctionSearcher struct{}

func (s *FunctionSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"cloudrun_functions",
		wf,
		cfg,
		q,
		gc.ListCloudRunFunctions,
		func(wf *aw.Workflow, gcrf gc.CloudRunFunction) {
			crf := FromGCloudCloudRunFunction(&gcrf)
			resource.NewItem(wf, cfg, crf, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
