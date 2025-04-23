package orchestrator

import (
	"errors"
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

var ErrNoActiveConfig = errors.New("no active gcloud config found")

const configCacheKey = "gcloud_config"

func getActiveConfig(ctx *Context, overrideCache bool) *gcloud.Config {
	wf := ctx.Workflow
	if !overrideCache && wf.Cache.Exists(configCacheKey) {
		log.Println("LOG: loading active config from cache")
		var cached gcloud.Config
		if err := wf.Cache.LoadJSON(configCacheKey, &cached); err == nil {
			return &cached
		}
	}

	log.Println("LOG: loading active config from gcloud")

	// load active config from gcloud
	active := gcloud.GetActiveConfig(wf)
	if active == nil {
		return nil
	}

	if err := wf.Cache.StoreJSON(configCacheKey, active); err != nil {
		wf.NewItem("Error caching active config").
			Subtitle("Please check the logs for more details").
			Icon(aw.IconError).
			Valid(false)
		wf.SendFeedback()
		return nil
	}

	return active
}
