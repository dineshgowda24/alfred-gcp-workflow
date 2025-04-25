package orchestrator

import (
	"log"
)

var _ Handler = (*FallbackHandler)(nil)

type FallbackHandler struct{}

func (h *FallbackHandler) Handle(ctx *Context) error {
	log.Println("LOG: FallbackHandler started")

	wf := ctx.Workflow
	config := ctx.ActiveConfig

	for _, svc := range ctx.Services {
		url, err := svc.Url(config)
		if err != nil {
			log.Printf("LOG: Error generating URL for service %s: %v\n", svc.Name, err)
		}

		wf.NewItem(svc.Name).
			Subtitle(svc.Subtitle()).
			Autocomplete(svc.ID).
			Arg(url).
			Icon(svc.Icon()).
			Valid(true)
	}

	wf.Filter(ctx.Args.Query)
	wf.SendFeedback()

	log.Println("LOG: FallbackHandler complete")
	return nil
}
