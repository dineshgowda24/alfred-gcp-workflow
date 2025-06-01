package orchestrator

import (
	"fmt"

	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/log"
)

var _ Handler = (*ServiceHandler)(nil)

type ServiceHandler struct{}

func (s *ServiceHandler) Handle(ctx *Context) error {
	service := ctx.ParsedQuery.Service
	log.Debugf("service handler started for: %s", service.Name)

	if len(service.SubServices) == 0 {
		s.handleNoSubservices(ctx, service)
		return nil
	}

	log.Debugf("listing subservices for: %s", service.Name)

	for i := range service.SubServices {
		s.addSubservice(ctx, &service.SubServices[i])
	}

	s.send(ctx)
	log.Debugf("service handler completed for: %s", service.Name)
	return nil
}

func (s *ServiceHandler) handleNoSubservices(ctx *Context, svc *services.Service) {
	wf := ctx.Workflow
	url, _ := svc.Url(ctx.ActiveConfig)

	wf.NewItem(svc.Autocomplete()).
		Subtitle(svc.Subtitle(nil)).
		Match(svc.Match()).
		Autocomplete(buildAutocomplete(ctx, svc)).
		Arg(url).
		Icon(svc.Icon()).
		Valid(true)

	addContributingItem(wf, fmt.Sprintf("%s has no sub-services (yet)", svc.Name))
	s.send(ctx)
}

func (s *ServiceHandler) addSubservice(ctx *Context, svc *services.Service) {
	wf := ctx.Workflow
	url, err := svc.Url(ctx.ActiveConfig)
	if err != nil {
		return
	}

	wf.NewItem(svc.Title()).
		Subtitle(svc.Subtitle(ctx.SearchRegistry)).
		Match(svc.Match()).
		Autocomplete(buildAutocomplete(ctx, svc)).
		Arg(url).
		Icon(svc.Icon()).
		Valid(true)
}

func (s *ServiceHandler) send(ctx *Context) {
	wf := ctx.Workflow

	wf.Filter(ctx.ParsedQuery.RemainingQuery)
	if wf.IsEmpty() {
		emptyResultItem(wf, "No matching subservices found")
	}

	wf.SendFeedback()
}
