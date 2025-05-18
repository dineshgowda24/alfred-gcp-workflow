package orchestrator

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/log"
)

var _ Handler = (*SubServiceHandler)(nil)

type SubServiceHandler struct{}

func (h *SubServiceHandler) Handle(ctx *Context) error {
	log.Debug("subservice handler started")

	query := ctx.ParsedQuery
	service := query.Service
	sub := query.SubService

	searcher := ctx.GetSearcher(service, sub)
	if searcher != nil {
		return h.handleSearcher(ctx, searcher)
	}

	h.addFallbackItem(ctx)
	log.Debug("subservice handler completed")
	ctx.Workflow.SendFeedback()
	return nil
}

func (h *SubServiceHandler) handleSearcher(ctx *Context, searcher searchers.Searcher) error {
	log.Debugf("running searcher: %s", ctx.ParsedQuery.SubService.Name)
	sub := ctx.ParsedQuery.SubService
	wf := ctx.Workflow
	if err := searcher.Search(wf, sub, ctx.ActiveConfig, ctx.ParsedQuery); err != nil {
		return err
	}

	wf.Filter(ctx.ParsedQuery.RemainingQuery)
	if wf.IsEmpty() {
		wf.NewItem("No results found 🤷").
			Subtitle("May be try a different query?").
			Icon(aw.IconTrash).
			Valid(false)
	}
	wf.SendFeedback()

	log.Debug("subservice searcher completed")
	return nil
}

func (h *SubServiceHandler) addFallbackItem(ctx *Context) {
	sub := ctx.ParsedQuery.SubService
	wf := ctx.Workflow

	url, err := sub.Url(ctx.ActiveConfig)
	if err != nil {
		log.Errorf("error getting URL for subservice %s: %v", sub.Name, err)
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
