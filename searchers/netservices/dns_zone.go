package netservices

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type DNSZoneSearcher struct{}

func (s *DNSZoneSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"netservices_dns_zones",
		wf,
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gc.DNSZone) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *DNSZoneSearcher) fetch(config *gc.Config) ([]gc.DNSZone, error) {
	return gc.ListDNSZones(config)
}

func (s *DNSZoneSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.DNSZone) {
	zone := DNSZoneFromGCloud(&entity)
	wf.NewItem(zone.Title()).
		Subtitle(zone.Subtitle()).
		Arg(zone.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
