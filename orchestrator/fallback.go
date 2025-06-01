package orchestrator

import (
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/log"
)

var _ Handler = (*FallbackHandler)(nil)

type FallbackHandler struct{}

func (f *FallbackHandler) Handle(ctx *Context) error {
	log.Debug("fallback handler started")

	for i := range ctx.Services {
		f.addService(ctx, &ctx.Services[i])
	}
	f.send(ctx)

	log.Debug("fallback handler completed")
	return nil
}

func (f *FallbackHandler) addService(ctx *Context, s *services.Service) {
	url, err := s.Url(ctx.ActiveConfig)
	if err != nil {
		return
	}

	ctx.Workflow.NewItem(s.ID).
		Subtitle(s.Subtitle(nil)).
		Match(s.Match()).
		Autocomplete(buildAutocomplete(ctx, s)).
		Arg(url).
		Icon(s.Icon()).
		Valid(true)
}

func (f *FallbackHandler) send(ctx *Context) {
	wf := ctx.Workflow
	wf.Filter(ctx.ParsedQuery.RemainingQuery)

	if wf.IsEmpty() {
		emptyResultItem(wf, "No matching services found")
	}

	wf.SendFeedback()
}
