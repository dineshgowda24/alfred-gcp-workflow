package orchestrator

import (
	"log"
	"strings"
)

var _ Handler = (*FallbackHandler)(nil)

type FallbackHandler struct{}

func (h *FallbackHandler) Handle(ctx *Context) error {
	log.Println("LOG: FallbackHandler started")

	wf := ctx.Workflow
	config := ctx.ActiveConfig
	query := ctx.Args.Query
	remaining := ctx.ParsedQuery.RemainingQuery
	isConfigQuery := ctx.ParsedQuery.IsConfigQuery

	for _, svc := range ctx.Services {
		url, err := svc.Url(config)
		if err != nil {
			log.Printf("LOG: Error generating URL for service %s: %v\n", svc.Name, err)
			continue
		}

		wf.NewItem(svc.ID).
			Subtitle(svc.Subtitle(nil)).
			Autocomplete(buildAutocompleteForFallback(query, remaining, svc.ID, isConfigQuery)).
			Arg(url).
			Icon(svc.Icon()).
			Valid(true)
	}

	log.Println("LOG: remaining query:", remaining)
	wf.Filter(remaining)
	wf.SendFeedback()

	log.Println("LOG: FallbackHandler complete")
	return nil
}

func buildAutocompleteForFallback(query, remaining, serviceID string, isConfigQuery bool) string {
	if !isConfigQuery {
		return serviceID
	}

	// If query includes partial config, replace or append serviceID properly
	stripped := strings.TrimSuffix(query, remaining)
	stripped = strings.TrimRight(stripped, " ")
	if stripped != "" {
		return stripped + " " + serviceID
	}
	return query + " " + serviceID
}
