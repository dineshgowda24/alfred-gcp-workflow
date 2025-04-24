package workflow

import (
	"os"
	"os/exec"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

const bgJobName = "alfred-gcp-workflow-bg"

type (
	FetchFunc[T any]  func(config *gcloud.Config) ([]T, error)
	RenderFunc[T any] func(wf *aw.Workflow, entity T)
)

type RenderRequest[T any] struct {
	Wf     *aw.Workflow
	Config *gcloud.Config
	Args   *SearchArgs
	Key    string
	Fetch  FetchFunc[T]
	Render RenderFunc[T]
}

func ResolveAndRender[T any](req RenderRequest[T]) error {
	if req.Args.RebuildCache {
		return fetchAndStore(req)
	}

	if tryRenderFromCache(req) {
		return nil
	}

	showLoadingState(req)
	triggerBgJob(req)

	req.Wf.Rerun(0.5)
	return nil
}

func fetchAndStore[T any](req RenderRequest[T]) error {
	data, err := req.Fetch(req.Config)
	if err != nil {
		return err
	}
	return req.Wf.Cache.StoreJSON(req.Config.CacheKey(req.Key), data)
}

func tryRenderFromCache[T any](req RenderRequest[T]) bool {
	cacheKey := req.Config.CacheKey(req.Key)
	maxAge := 5 * time.Minute

	if req.Wf.Cache.Expired(cacheKey, maxAge) {
		return false
	}

	var results []T
	if err := req.Wf.Cache.LoadJSON(cacheKey, &results); err != nil {
		return false
	}

	for _, item := range results {
		req.Render(req.Wf, item)
	}
	return true
}

func showLoadingState[T any](req RenderRequest[T]) {
	req.Wf.NewItem("Loading...").
		Subtitle("Fetching GCP resources").
		Icon(aw.IconInfo).
		Valid(false)
}

func triggerBgJob[T any](req RenderRequest[T]) {
	if req.Wf.IsRunning(bgJobName) {
		return
	}
	cmd := exec.Command(os.Args[0], "--query="+req.Args.Query, "--rebuild-cache")
	_ = req.Wf.RunInBackground(bgJobName, cmd) // handle/log error if needed
}
