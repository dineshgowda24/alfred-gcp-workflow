package vpc

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type RouteSearcher struct{}

func (r *RouteSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"vpc_routes",
		wf,
		config,
		pq,
		r.fetch,
		func(wf *aw.Workflow, entity gc.VPCRoute) {
			r.render(wf, svc, config, entity)
		},
	))
}

func (r *RouteSearcher) fetch(config *gc.Config) ([]gc.VPCRoute, error) {
	return gc.ListVPCRoutes(config)
}

func (r *RouteSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.VPCRoute) {
	route := RouteFromGCloud(&entity)
	wf.NewItem(route.Title()).
		Subtitle(route.Subtitle()).
		Arg(route.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
