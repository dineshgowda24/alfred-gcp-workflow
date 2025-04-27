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

	configs, err := gcloud.GetAllConfigs()
	if err != nil {
		log.Println("LOG: error getting all configs:", err)
		return err
	}

	for _, config := range configs {
		h.addConfigItem(ctx, config)
	}

	h.send(ctx)
	log.Println("LOG: ConfigHandler complete")
	return nil
}

func (h *ConfigHandler) buildConfigAutocomplete(ctx *Context, config *gcloud.Config) string {
	query := ctx.Args.Query
	partial := ctx.ParsedQuery.PartialConfigQuery
	full := config.Name
	return strings.Replace(query, "@"+partial, "@"+full, 1)
}

func (h *ConfigHandler) addConfigItem(ctx *Context, config *gcloud.Config) {
	newQuery := h.buildConfigAutocomplete(ctx, config)
	ctx.Workflow.NewItem(config.Name).
		Subtitle(config.Project).
		Autocomplete(newQuery).
		Arg(newQuery).
		Icon(aw.IconAccount).
		Valid(true)
}

func (h *ConfigHandler) send(ctx *Context) {
	wf := ctx.Workflow
	wf.Filter(ctx.ParsedQuery.PartialConfigQuery)

	if wf.IsEmpty() {
		emptyResultItem(wf, "No matching configurations found")
	}

	wf.SendFeedback()
}
