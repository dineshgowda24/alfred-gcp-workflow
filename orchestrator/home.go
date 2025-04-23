package orchestrator

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type homeItem struct {
	Title    string
	Subtitle string
	Arg      string
	Icon     *aw.Icon
	Valid    bool
}

var _ Handler = (*HomeHandler)(nil)

type HomeHandler struct{}

func (h *HomeHandler) Handle(ctx *Context) error {
	items := []homeItem{
		{
			Title:    "🔍 Search GCP Services",
			Subtitle: "Try keywords like: sql, redis, storage",
			Icon:     &aw.Icon{Value: "/Users/dinesh.chikkanna/personal/alfred-gcp-workflow/images/gcp.png"},
			Valid:    false,
		},
		{
			Title:    "🧭 Open Console",
			Subtitle: "Go to " + gcloud.ConsoleURL,
			Arg:      gcloud.ConsoleURL,
			Icon:     aw.IconWeb,
			Valid:    true,
		},
		{
			Title:    "🩺 GCP Health Status",
			Subtitle: "Check service availability across regions",
			Arg:      gcloud.HealthStatusURL,
			Icon:     &aw.Icon{Value: "/Users/dinesh.chikkanna/personal/alfred-gcp-workflow/images/heartbeat.png"},
			Valid:    true,
		},
	}

	if ctx.ActiveConfig != nil {
		items = append(items, homeItem{
			Title:    "☁️ Active gcloud config: " + ctx.ActiveConfig.Name,
			Subtitle: "Currently selected gcloud configuration",
			Valid:    false,
		})
	} else {
		items = append(items, homeItem{
			Title:    "⚠️ No active gcloud config found",
			Subtitle: "Click to open gcloud quickstart guide",
			Icon:     aw.IconError,
			Arg:      gcloud.QuickStartURL,
			Valid:    false,
		})
	}

	for _, item := range items {
		it := ctx.Workflow.NewItem(item.Title).
			Subtitle(item.Subtitle).
			Valid(item.Valid)
		if item.Arg != "" {
			it.Arg(item.Arg)
		}
		if item.Icon != nil {
			it.Icon(item.Icon)
		}
	}

	ctx.SendFeedback()
	return nil
}
