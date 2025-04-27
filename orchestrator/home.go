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
			Title:    "Search Services",
			Subtitle: "🔍 Try keywords like: sql, memorystore, storage",
			Icon:     aw.IconWorkflow,
			Valid:    false,
		},
		{
			Title:    "Open Cloud Console",
			Subtitle: "🌐 Launch: " + gcloud.ConsoleURL,
			Arg:      gcloud.ConsoleURL,
			Icon:     aw.IconHome,
			Valid:    true,
		},
		{
			Title:    "Health Dashboard",
			Subtitle: "🏥 View service availability by region",
			Arg:      gcloud.HealthStatusURL,
			Icon:     &aw.Icon{Value: "/Users/dinesh.chikkanna/personal/alfred-gcp-workflow/images/heartbeat.png"},
			Valid:    true,
		},
		{
			Title:    "Use @ to override config",
			Subtitle: "⚙️ Example: @default",
			Icon:     aw.IconAccount,
			Valid:    false,
		},
	}

	if ctx.ActiveConfig != nil {
		items = append(items, homeItem{
			Title:    "Active gcloud config: " + ctx.ActiveConfig.Name,
			Subtitle: "🔐 Currently selected configuration",
			Icon:     aw.IconSettings,
			Valid:    false,
		})
	} else {
		items = append(items, homeItem{
			Title:    "No active gcloud config found",
			Subtitle: "⚠️ Click to open the gcloud quickstart guide",
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
