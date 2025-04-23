package orchestrator

import (
	"log"

	aw "github.com/deanishe/awgo"
)

var _ Handler = (*ServiceHandler)(nil)

type ServiceHandler struct{}

func (h *ServiceHandler) Handle(ctx *Context) error {
	parsedQuery := ctx.ParsedQuery
	wf := ctx.Workflow
	active := ctx.ActiveConfig
	parent := parsedQuery.Service

	log.Println("LOG: service found:", parent.Name)
	log.Println("LOG: listing subservices for service:", parent.Name)

	for _, child := range parent.SubServices {
		url, err := child.Url(active)
		if err != nil {
			log.Println("LOG: error rendering subservice URL:", err)
		}

		subtitle := child.Description
		searcher := ctx.GetSearcher(parent, &child)
		if searcher != nil {
			subtitle = "🔍 " + subtitle
		}

		wf.NewItem(child.Name).
			Subtitle(subtitle).
			Autocomplete(parent.ID + " " + child.ID).
			Icon(child.Icon()).
			Arg(url).
			Valid(true)
	}

	if len(parsedQuery.Service.SubServices) == 0 {
		log.Println("LOG: no subservices found for service:", parsedQuery.Service.Name)
		url, err := parent.Url(active)
		if err != nil {
			log.Println("LOG: error rendering URL:", err)
		}
		wf.NewItem(parent.Name).
			Autocomplete(parent.ID).
			Subtitle(parent.Description).
			Arg(url).
			Icon(parent.Icon()).
			Valid(true)

		wf.NewFileItem(parent.Name + " has no sub-services (yet)").
			Subtitle("Open contributing guide to add them").
			Arg("https://github.com/dineshgowda24/alfred-gcp-workflow").
			Icon(aw.IconNote).
			Valid(true)

	}

	wf.SendFeedback()
	return nil
}
