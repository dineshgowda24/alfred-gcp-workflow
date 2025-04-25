package workflow

import (
	"log"
	"os"
	"os/exec"
	"time"

	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
)

const (
	bgJobName   = "alfred-gcp-workflow-background-job"
	cacheMaxAge = 5 * time.Minute
)

type (
	FetchFunc[T any]  func(config *gc.Config) ([]T, error)
	RenderFunc[T any] func(wf *aw.Workflow, entity T)
)

type RenderRequest[T any] struct {
	Key string

	Wf     *aw.Workflow
	Config *gc.Config

	Fetch  FetchFunc[T]
	Render RenderFunc[T]
	Pq     *parser.Result
}

func NewRenderRequest[T any](
	key string, wf *aw.Workflow, cfg *gc.Config, pq *parser.Result, fetch FetchFunc[T], render RenderFunc[T],
) RenderRequest[T] {
	return RenderRequest[T]{
		Key:    key,
		Wf:     wf,
		Config: cfg,
		Pq:     pq,
		Fetch:  fetch,
		Render: render,
	}
}

func ResolveAndRender[T any](r RenderRequest[T]) error {
	if r.Pq.SearchArgs.RebuildCache {
		return store(r)
	}

	if ok := renderCache(r); ok {
		return nil
	}

	loadingItem(r)
	spawnBgJob(r)
	r.Wf.Rerun(0.2)
	return nil
}

func store[T any](r RenderRequest[T]) error {
	data, err := r.Fetch(r.Config)
	if err != nil {
		return err
	}
	return r.Wf.Cache.StoreJSON(r.Config.CacheKey(r.Key), data)
}

func renderCache[T any](r RenderRequest[T]) bool {
	cacheKey := r.Config.CacheKey(r.Key)

	if r.Wf.Cache.Expired(cacheKey, cacheMaxAge) {
		return false
	}

	log.Println("LOG: cache hit for", r.Key)
	var results []T
	if err := r.Wf.Cache.LoadJSON(cacheKey, &results); err != nil {
		return false
	}

	for _, item := range results {
		r.Render(r.Wf, item)
	}
	return true
}

func loadingItem[T any](r RenderRequest[T]) {
	r.Wf.NewItem("🔄 Loading GCP resources...").
		Subtitle("Fetching data and preparing results. Please wait 🙏").
		Icon(aw.IconSync).
		Valid(false)
}

func spawnBgJob[T any](r RenderRequest[T]) {
	if r.Wf.IsRunning(bgJobName) {
		return
	}
	cmd := exec.Command(os.Args[0], "--query="+r.Pq.SearchArgs.Query, "--rebuild-cache")
	if err := r.Wf.RunInBackground(bgJobName, cmd); err != nil {
		panic(err) // should never happen
	}
}
