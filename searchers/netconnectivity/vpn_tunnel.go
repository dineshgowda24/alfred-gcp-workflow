package netconnectivity

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type VPNTunnelSearcher struct{}

func (s *VPNTunnelSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"netconnectivity_vpn_tunnels",
		wf,
		cfg,
		q,
		gc.ListVPNTunnels,
		func(wf *aw.Workflow, gvt gc.VPNTunnel) {
			vt := FromGCloudVPNTunnel(&gvt)
			resource.NewItem(wf, cfg, vt, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
