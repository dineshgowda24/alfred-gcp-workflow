package netconnectivity

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type VPNTunnelSearcher struct{}

func (s *VPNTunnelSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"networkconnectivity_vpn_tunnels",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.VPNTunnel) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *VPNTunnelSearcher) fetch(config *gc.Config) ([]gc.VPNTunnel, error) {
	return gc.ListVPNTunnels(config)
}

func (s *VPNTunnelSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.VPNTunnel) {
	ins := VPNTunnelFromGCloud(&entity)
	wf.NewItem(ins.Name).
		Subtitle(ins.Subtitle()).
		Arg(ins.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
