package orchestrator

import (
	"fmt"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/config"
)

type IntentHandler struct{}

func (h *IntentHandler) Handle(ctx *Context) error {
	wf := ctx.Workflow
	q := ctx.ParsedQuery

	if q.Intent != "gcloud-path" {
		return nil
	}

	path := strings.TrimSpace(q.IntentValue)
	if path == "" {
		wf.NewItem("One step away from unlocking the realm 🧩").
			Subtitle("Usage: gcloud-path /path/to/gcloud").
			Arg("gcloud-path ").
			Autocomplete("gcloud-path ").
			Icon(aw.IconSettings).
			Valid(false)
		wf.SendFeedback()
		return nil
	}

	bin := gcloud.NormalizeGCloudPath(path)
	info, err := gcloud.GetGCloudInfo(bin)
	if err != nil {
		return err
	}

	cfg := &config.ConfigFile{
		GCloudPath:       bin,
		GCloudConfigPath: info.Config.Paths.GlobalConfigDir,
	}
	config.UpdateConfigFile(wf, cfg)

	wf.NewItem("Relm unlocked! 🗝️").
		Subtitle(fmt.Sprintf("gcloud path captured ✅ | Version: %s", info.Installation.Components.Core)).
		Icon(aw.IconInfo).
		Autocomplete("").
		Arg("").
		Valid(false)

	wf.SendFeedback()
	return nil
}
