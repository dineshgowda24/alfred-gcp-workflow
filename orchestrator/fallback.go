package orchestrator

import "log"

var _ Handler = (*FallbackHandler)(nil)

type FallbackHandler struct{}

func (h *FallbackHandler) Handle(ctx *Context) error {
	log.Println("LOG: running fallback handler")
	wf := ctx.Workflow
	config := ctx.ActiveConfig

	for _, svc := range ctx.Services {
		url, err := svc.Url(config)
		if err != nil {
			log.Println("LOG: error rendering URL:", err)
		}

		wf.NewItem(svc.Name).
			Autocomplete(svc.ID).
			Subtitle(svc.Subtitle()).
			Arg(url).
			Icon(svc.Icon()).
			Valid(true)
	}
	log.Println("LOG: running fallback handler complete")

	wf.Filter(ctx.RawQuery)
	wf.SendFeedback()
	return nil
}
