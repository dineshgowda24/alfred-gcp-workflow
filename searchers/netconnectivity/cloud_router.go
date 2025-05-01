package netconnectivity

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type CloudRouterSearcher struct{}

func (s *CloudRouterSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"netconnectivity_cloud_routers",
		wf,
		cfg,
		q,
		gc.ListCloudRouters,
		func(wf *aw.Workflow, gcr gc.CloudRouter) {
			cr := FromGCloudCloudRouter(&gcr)
			resource.NewItem(wf, cfg, cr, svc.Icon())
		},
	)

	return builder.Build()
}
