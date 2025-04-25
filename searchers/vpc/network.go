package vpc

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type NetworkSearcher struct{}

func (s *NetworkSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"vpc_networks",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.VPCNetwork) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *NetworkSearcher) fetch(config *gc.Config) ([]gc.VPCNetwork, error) {
	return gc.ListVPCNetworks(config)
}

func (s *NetworkSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.VPCNetwork) {
	network := NetworkFromGCloud(&entity)
	wf.NewItem(network.Title()).
		Subtitle(network.Subtitle()).
		Arg(network.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
