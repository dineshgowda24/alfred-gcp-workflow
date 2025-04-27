package orchestrator

import (
	"log"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

var _ Handler = (*ConfigHandler)(nil)

type ConfigHandler struct{}

func (h *ConfigHandler) Handle(ctx *Context) error {
	log.Println("LOG: ConfigHandler started")

	wf := ctx.Workflow
	query := ctx.Args.Query
	partial := ctx.ParsedQuery.PartialConfigQuery

	configs, err := gcloud.GetAllConfigs()
	if err != nil {
		log.Println("LOG: error getting all configs:", err)
		return err
	}

	for _, config := range configs {
		newQuery := buildConfigAutocomplete(query, partial, config.Name)
		wf.NewItem(config.Name).
			Subtitle(config.Project).
			Autocomplete(newQuery).
			Arg(newQuery).
			Icon(aw.IconAccount).
			Valid(true)
	}

	wf.Filter(partial)

	if wf.IsEmpty() {
		wf.NewItem("No matching configurations found").
			Subtitle("Try a different query with @").
			Icon(aw.IconNote).
			Valid(false)
	}

	wf.SendFeedback()

	log.Println("LOG: ConfigHandler complete")
	return nil
}

func buildConfigAutocomplete(query, partial, full string) string {
	return strings.Replace(query, "@"+partial, "@"+full, 1)
}
