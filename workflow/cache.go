package workflow

import (
	"os"
	"os/exec"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

const bgName = "alfred-gcp-workflow-bg"

type (
	FetchFunc[T any]  func(config *gcloud.Config) ([]T, error)
	RenderFunc[T any] func(wf *aw.Workflow, entity T)
)

func LoadFromCache[T any](
	wf *aw.Workflow,
	cacheKey string,
	config *gcloud.Config,
	args *SearchArgs,
	fetcher FetchFunc[T],
	render RenderFunc[T],
) error {
	if args.RebuildCache {
		data, err := fetcher(config)
		if err != nil {
			return err
		}
		return wf.Cache.StoreJSON(cacheKey, data)
	}

	var results []T
	maxAge := 5 * time.Minute
	if !wf.Cache.Expired(cacheKey, maxAge) {
		if err := wf.Cache.LoadJSON(cacheKey, &results); err == nil {
			for _, item := range results {
				render(wf, item)
			}
			return nil
		}
	}

	// Cache miss: show loading + trigger fetch
	wf.NewItem("Loading...").
		Subtitle("Fetching GCP resources").
		Icon(aw.IconInfo).
		Valid(false)

	if !wf.IsRunning(bgName) {
		cmd := exec.Command(os.Args[0], "--query="+args.Query, "--rebuild-cache")
		if err := wf.RunInBackground(bgName, cmd); err != nil {
			panic(err)
		}
	}

	wf.Rerun(0.5)
	return nil
}
