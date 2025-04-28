package netconnectivity

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type VPNGatewaySearcher struct{}

func (s *VPNGatewaySearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"netconnectivity_vpn_gateways",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.VPNGateway) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *VPNGatewaySearcher) fetch(config *gc.Config) ([]gc.VPNGateway, error) {
	return gc.ListVPNGateways(config)
}

func (s *VPNGatewaySearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.VPNGateway) {
	vpnGateway := VPNGatewayFromGCloud(&entity)
	wf.NewItem(vpnGateway.Name).
		Subtitle(vpnGateway.Subtitle()).
		Arg(vpnGateway.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
