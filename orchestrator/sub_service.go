package orchestrator

import (
	"log"

	aw "github.com/deanishe/awgo"
)

var _ Handler = (*SubServiceHandler)(nil)

type SubServiceHandler struct{}

func (h *SubServiceHandler) Handle(ctx *Context) error {
	log.Println("LOG: run subservice handler")
	query := ctx.ParsedQuery
	wf := ctx.Workflow
	active := ctx.ActiveConfig
	parent := query.Service
	child := query.SubService

	searcher := ctx.GetSearcher(parent, child)
	if searcher != nil {
		log.Println("LOG: executing searcher for subservice:", child.Name)
		err := searcher.Search(wf, child, active, query.RemainingQuery)
		if err != nil {
			return err
		}
		wf.Filter(query.RemainingQuery)

	} else {
		wf.NewItem(child.Name).
			Subtitle(child.Description).
			Autocomplete(query.Service.ID + " " + child.ID).
			Icon(child.Icon()).
			Arg(child.URL).
			Valid(true)

		wf.NewFileItem(child.Name + " has no searcher (yet)").
			Subtitle("Open contributing guide to add them").
			Arg("https://github.com/dineshgowda24/alfred-gcp-workflow").
			Icon(aw.IconNote).
			Valid(true)

	}

	log.Println("LOG: run subservice handler complete")
	wf.SendFeedback()
	return nil
}
