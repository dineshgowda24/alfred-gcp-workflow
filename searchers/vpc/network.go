package vpc

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type NetworkSearcher struct{}

func (s *NetworkSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, pq *parser.Result,
) error {
	return resource.NewBuilder(
		"vpc_networks",
		wf,
		cfg,
		pq,
		gc.ListVPCNetworks,
		func(wf *aw.Workflow, gvn gc.VPCNetwork) {
			nw := FromGCloudNetwork(&gvn)
			resource.NewItem(wf, cfg, nw, svc.Icon())
		},
	).Build()
}
