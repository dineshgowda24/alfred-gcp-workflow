package orchestrator

import (
	"fmt"
	"log"

	"github.com/dineshgowda24/alfred-gcp-workflow/searchers"
)

var _ Handler = (*SubServiceHandler)(nil)

type SubServiceHandler struct{}

func (h *SubServiceHandler) Handle(ctx *Context) error {
	log.Println("LOG: SubServiceHandler started")

	query := ctx.ParsedQuery
	service := query.Service
	sub := query.SubService

	searcher := ctx.GetSearcher(service, sub)
	if searcher != nil {
		return h.handleSearcher(ctx, searcher)
	}

	h.addFallbackItem(ctx)
	log.Println("LOG: SubServiceHandler completed")
	ctx.Workflow.SendFeedback()
	return nil
}

func (h *SubServiceHandler) handleSearcher(ctx *Context, searcher searchers.Searcher) error {
	sub := ctx.ParsedQuery.SubService
	log.Printf("LOG: Found searcher for subservice: %s", sub.Name)
	wf := ctx.Workflow
	if err := searcher.Search(wf, sub, ctx.ActiveConfig, ctx.ParsedQuery); err != nil {
		return err
	}

	wf.Filter(ctx.ParsedQuery.RemainingQuery)
	wf.SendFeedback()

	log.Println("LOG: Searcher flow complete")
	return nil
}

func (h *SubServiceHandler) addFallbackItem(ctx *Context) {
	sub := ctx.ParsedQuery.SubService
	wf := ctx.Workflow

	url, err := sub.Url(ctx.ActiveConfig)
	if err != nil {
		log.Printf("LOG: Error generating URL for subservice %s: %v", sub.Name, err)
		return
	}

	wf.NewItem(sub.Title()).
		Subtitle(sub.Subtitle(nil)).
		Autocomplete(buildAutocomplete(ctx, sub)).
		Icon(sub.Icon()).
		Arg(url).
		Valid(true)

	addContributingItem(wf, fmt.Sprintf("%s has no searcher (yet)", sub.Name))
}
