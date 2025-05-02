package cloudrun

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type ServiceSearcher struct{}

func (s *ServiceSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"cloudrun_services",
		wf,
		cfg,
		q,
		gc.ListCloudRunServices,
		func(wf *aw.Workflow, gcrs gc.CloudRunService) {
			crs := FromGCloudCloudRunService(&gcrs)
			resource.NewItem(wf, cfg, crs, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
