package orchestrator

import (
	"log"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

var _ Handler = (*RegionHandler)(nil)

type RegionHandler struct{}

func (h *RegionHandler) Handle(ctx *Context) error {
	log.Println("region handler started")

	regions := gcloud.GetAllRegions()
	for _, region := range regions {
		h.addRegion(ctx, &region)
	}

	h.send(ctx)
	log.Println("region handler completed")
	return nil
}

func (h *RegionHandler) addRegion(ctx *Context, region *gcloud.Region) {
	newQuery := h.buildConfigAutocomplete(ctx, region)
	ctx.Workflow.NewItem(region.Name).
		Subtitle(region.Location).
		Autocomplete(newQuery).
		Arg(newQuery).
		Match(region.Name + " " + region.Location).
		Icon(aw.IconNetwork).
		Valid(true)
}

func (h *RegionHandler) buildConfigAutocomplete(ctx *Context, region *gcloud.Region) string {
	query := ctx.Args.Query
	partial := ctx.ParsedQuery.PartialRegionQuery
	return strings.Replace(query, "$"+partial, "$"+region.Name, 1)
}

func (h *RegionHandler) send(ctx *Context) {
	wf := ctx.Workflow
	wf.Filter(ctx.ParsedQuery.PartialRegionQuery)

	if wf.IsEmpty() {
		emptyResultItem(wf, "No matching regions found")
	}

	wf.SendFeedback()
}
