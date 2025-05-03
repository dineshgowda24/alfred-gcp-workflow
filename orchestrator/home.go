package orchestrator

import (
	"fmt"
	"math"
	"path/filepath"
	"slices"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

const (
	MagicPrefix = "tools:"
)

type homeItem struct {
	Title        string
	Subtitle     string
	Arg          string
	Icon         *aw.Icon
	Valid        bool
	AutoComplete string
	SortPriority int
}

var _ Handler = (*HomeHandler)(nil)

type HomeHandler struct{}

func (h *HomeHandler) Handle(ctx *Context) error {
	items := []homeItem{
		{
			Title:        "Search Services",
			Subtitle:     "🔍 Try keywords like: sql, memorystore, storage",
			Icon:         aw.IconWorkflow,
			Valid:        false,
			SortPriority: 1,
		},
		{
			Title:        "Open Cloud Console",
			Subtitle:     "🌐 Launch: " + gcloud.ConsoleURL,
			Arg:          gcloud.ConsoleURL,
			Icon:         aw.IconHome,
			Valid:        true,
			SortPriority: 2,
		},
		{
			Title:        "Health Dashboard",
			Subtitle:     "🏥 View service availability by region",
			Arg:          gcloud.HealthStatusURL,
			Icon:         &aw.Icon{Value: filepath.Join(ctx.Workflow.Dir(), "images", "heartbeat.png")},
			Valid:        true,
			SortPriority: 3,
		},
		{
			Title:        "Dev Tools",
			Subtitle:     "🧹 Clear cache, view logs, or reset internal state",
			Arg:          MagicPrefix,
			AutoComplete: MagicPrefix,
			Icon:         aw.IconInfo,
			Valid:        false,
			SortPriority: math.MaxInt,
		},
	}

	if ctx.ActiveConfig != nil {
		items = append(items, homeItem{
			Title:        fmt.Sprintf("Using configuration \"%s\"", ctx.ActiveConfig.Name),
			Subtitle:     "🔩 Use @ to override gcloud config",
			Icon:         aw.IconAccount,
			Valid:        false,
			SortPriority: 4,
		})

		if ctx.ActiveConfig.Region != nil {
			items = append(items, homeItem{
				Title:        fmt.Sprintf("Using region \"%s\"", ctx.ActiveConfig.Region.Name),
				Subtitle:     "🗺️ Use `$` to override region",
				Icon:         aw.IconNetwork,
				Valid:        false,
				SortPriority: 5,
			})
		} else {
			items = append(items, homeItem{
				Title:        "No region selected",
				Subtitle:     "🗺️ Use `$` to override region",
				Icon:         aw.IconNetwork,
				Valid:        false,
				SortPriority: 6,
			})
		}

	} else {
		items = append(items, homeItem{
			Title:        "No active gcloud config found",
			Subtitle:     "⚠️ Click to open the gcloud quickstart guide",
			Icon:         aw.IconError,
			Arg:          gcloud.QuickStartURL,
			Valid:        false,
			SortPriority: 1,
		})
	}
	slices.SortFunc(items, itemSort)

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
		if item.AutoComplete != "" {
			it.Autocomplete(item.AutoComplete)
		}
	}

	ctx.SendFeedback()
	return nil
}

func itemSort(a, b homeItem) int {
	if a.SortPriority < b.SortPriority {
		return -1
	}
	return 1
}
