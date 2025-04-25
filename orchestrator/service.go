package orchestrator

import (
	"fmt"
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

var _ Handler = (*ServiceHandler)(nil)

type ServiceHandler struct{}

func (h *ServiceHandler) Handle(ctx *Context) error {
	service := ctx.ParsedQuery.Service
	wf := ctx.Workflow
	active := ctx.ActiveConfig

	log.Printf("LOG: Handling service: %s\n", service.Name)

	if len(service.SubServices) == 0 {
		log.Printf("LOG: No subservices for service: %s\n", service.Name)

		url, err := service.Url(active)
		if err != nil {
			log.Println("LOG: Error generating service URL:", err)
		}

		wf.NewItem(service.Autocomplete()).
			Subtitle(service.Subtitle()).
			Autocomplete(service.Autocomplete()).
			Arg(url).
			Icon(service.Icon()).
			Valid(true)

		addMissingSubServiceItem(wf, service)

		wf.SendFeedback()
		return nil
	}

	log.Printf("LOG: Listing %d subservices for: %s\n", len(service.SubServices), service.Name)

	for _, child := range service.SubServices {
		url, err := child.Url(active)
		if err != nil {
			log.Printf("LOG: Error generating URL for subservice %s: %v\n", child.Name, err)
			continue
		}

		subtitle := child.Subtitle()
		if ctx.SearchRegistry.Exists(service, &child) {
			subtitle = "🔍⚡️ " + subtitle
		}

		wf.NewItem(child.Title()).
			Subtitle(subtitle).
			Autocomplete(child.Autocomplete()).
			Icon(child.Icon()).
			Arg(url).
			Valid(true)
	}

	wf.Filter(ctx.ParsedQuery.RemainingQuery)
	wf.SendFeedback()
	return nil
}

func addMissingSubServiceItem(wf *aw.Workflow, service *services.Service) {
	wf.NewItem(fmt.Sprintf("%s has no sub-services (yet)", service.Name)).
		Subtitle("Open contributing guide to add them").
		Arg("https://github.com/dineshgowda24/alfred-gcp-workflow").
		Icon(aw.IconNote).
		Valid(true)
}
