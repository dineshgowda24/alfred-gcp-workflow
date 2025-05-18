package orchestrator

import (
	"fmt"

	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/log"
)

var _ Handler = (*ServiceHandler)(nil)

type ServiceHandler struct{}

func (h *ServiceHandler) Handle(ctx *Context) error {
	service := ctx.ParsedQuery.Service
	log.Debugf("service handler started for: %s", service.Name)

	if len(service.SubServices) == 0 {
		h.handleNoSubservices(ctx, service)
		return nil
	}

	log.Debugf("listing subservices for: %s", service.Name)

	for i := range service.SubServices {
		h.addSubservice(ctx, &service.SubServices[i])
	}

	h.send(ctx)
	log.Debugf("service handler completed for: %s", service.Name)
	return nil
}

func (h *ServiceHandler) handleNoSubservices(ctx *Context, s *services.Service) {
	wf := ctx.Workflow
	url, _ := s.Url(ctx.ActiveConfig)

	wf.NewItem(s.Autocomplete()).
		Subtitle(s.Subtitle(nil)).
		Match(s.Match()).
		Autocomplete(buildAutocomplete(ctx, s)).
		Arg(url).
		Icon(s.Icon()).
		Valid(true)

	addContributingItem(wf, fmt.Sprintf("%s has no sub-services (yet)", s.Name))
	h.send(ctx)
}

func (h *ServiceHandler) addSubservice(ctx *Context, s *services.Service) {
	wf := ctx.Workflow
	url, err := s.Url(ctx.ActiveConfig)
	if err != nil {
		return
	}

	wf.NewItem(s.Title()).
		Subtitle(s.Subtitle(ctx.SearchRegistry)).
		Match(s.Match()).
		Autocomplete(buildAutocomplete(ctx, s)).
		Arg(url).
		Icon(s.Icon()).
		Valid(true)
}

func (h *ServiceHandler) send(ctx *Context) {
	wf := ctx.Workflow

	wf.Filter(ctx.ParsedQuery.RemainingQuery)
	if wf.IsEmpty() {
		emptyResultItem(wf, "No matching subservices found")
	}

	wf.SendFeedback()
}
