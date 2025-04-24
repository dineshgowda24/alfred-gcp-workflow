package orchestrator

import (
	"fmt"
	"log"

	aw "github.com/deanishe/awgo"
)

var _ Handler = (*SubServiceHandler)(nil)

type SubServiceHandler struct{}

func (h *SubServiceHandler) Handle(ctx *Context) error {
	log.Println("LOG: SubServiceHandler started")

	query := ctx.ParsedQuery
	parent := query.Service
	child := query.SubService
	searcher := ctx.GetSearcher(parent, child)

	if searcher != nil {
		log.Printf("LOG: Found searcher for subservice: %s\n", child.Name)
		if err := searcher.Search(ctx.Workflow, child, ctx.ActiveConfig, *ctx.Args); err != nil {
			return err
		}
		ctx.Workflow.Filter(query.RemainingQuery)
	} else {
		log.Printf("LOG: No searcher found for subservice: %s\n", child.Name)

		ctx.Workflow.NewItem(child.Name).
			Subtitle(child.Description).
			Autocomplete(fmt.Sprintf("%s %s", parent.ID, child.ID)).
			Icon(child.Icon()).
			Arg(child.URL).
			Valid(true)

		ctx.Workflow.NewFileItem(fmt.Sprintf("%s has no searcher (yet)", child.Name)).
			Subtitle("Open contributing guide to add them").
			Arg("https://github.com/dineshgowda24/alfred-gcp-workflow").
			Icon(aw.IconNote).
			Valid(true)
	}

	log.Println("LOG: SubServiceHandler completed")
	ctx.Workflow.SendFeedback()
	return nil
}
