package workflow

import (
	"log"
	"os"
	"os/exec"
	"time"

	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/env"
)

const (
	bgJobName               = "alfred-gcp-workflow-background-job"
	spinnerIdxKey           = "spinner-index"
	defaultCacheTTLDuration = 7 * 24 * time.Hour
)

var (
	cacheTTLDuration = env.CacheTTLDuration(defaultCacheTTLDuration)
	spinnerFrames    = []string{"⏳", "⌛", "⏳", "⌛"}
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

	if r.Wf.Cache.Expired(cacheKey, cacheTTLDuration) {
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

func spawnBgJob[T any](r RenderRequest[T]) {
	if r.Wf.IsRunning(bgJobName) {
		return
	}
	cmd := exec.Command(os.Args[0], "--query="+r.Pq.SearchArgs.Query, "--rebuild-cache")
	if err := r.Wf.RunInBackground(bgJobName, cmd); err != nil {
		panic(err) // should never happen
	}
}

func loadingItem[T any](r RenderRequest[T]) {
	frame := nextSpinnerFrame(r.Wf)
	r.Wf.NewItem(frame + " Fetching resources").
		Subtitle("Hang tight 🙏").
		Icon(aw.IconInfo).
		Valid(false)
}

func nextSpinnerFrame(wf *aw.Workflow) string {
	var idx int
	if wf.Cache.Exists(spinnerIdxKey) {
		err := wf.Cache.LoadJSON(spinnerIdxKey, &idx)
		if err != nil {
			log.Println("LOG: error loading index from cache:", err)
		}
	}

	spinner := spinnerFrames[idx%len(spinnerFrames)]

	nextIdx := (idx + 1) % len(spinnerFrames)
	err := wf.Cache.StoreJSON(spinnerIdxKey, nextIdx)
	if err != nil {
		log.Println("LOG: error storing index to cache:", err)
	}
	return spinner
}
