package workflow

import (
	"log"
	"os"
	"os/exec"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/env"
)

const (
	bgJobName               = "alfred-gcp-workflow-background-job"
	spinnerIdxKey           = "spinner-index"
	defaultCacheTTLDuration = 7 * 24 * time.Hour
	defaultErrorCacheTTL    = 1 * time.Minute
)

var (
	cacheTTLDuration = env.CacheTTLDuration(defaultCacheTTLDuration)
	spinnerFrames    = []string{
		"🕐", "🕑", "🕒", "🕓", "🕔", "🕕",
		"🕖", "🕗", "🕘", "🕙", "🕚", "🕛",
	}
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
		r.Wf.Cache.Store(errorCacheKey(r.Key, r.Config), []byte(err.Error()))
		return err
	}
	return r.Wf.Cache.StoreJSON(r.Config.CacheKey(r.Key), data)
}

func errorCacheKey(key string, c *gcloud.Config) string {
	return c.CacheKey(key) + "_gcloud_error"
}

func renderCache[T any](r RenderRequest[T]) bool {
	if renderGCloudError(r) {
		return true
	}

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

func renderGCloudError[T any](r RenderRequest[T]) bool {
	cacheKey := errorCacheKey(r.Key, r.Config)
	if !r.Wf.Cache.Expired(cacheKey, defaultErrorCacheTTL) {
		msg, err := r.Wf.Cache.Load(cacheKey)
		if err != nil {
			msg = []byte(err.Error())
		}

		r.Wf.NewItem(string(msg)).
			Subtitle("gcloud cmd failed to execute").
			Icon(aw.IconError).
			Valid(false)
		r.Wf.SendFeedback()
		return true
	}

	return false
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
	var currIdx int
	if wf.Cache.Exists(spinnerIdxKey) {
		err := wf.Cache.LoadJSON(spinnerIdxKey, &currIdx)
		if err != nil {
			log.Println("LOG: error loading index from cache:", err)
		}
	}

	spinner := spinnerFrames[currIdx%len(spinnerFrames)]

	idx := (currIdx + 1) % len(spinnerFrames)
	err := wf.Cache.StoreJSON(spinnerIdxKey, idx)
	if err != nil {
		log.Println("LOG: error storing index to cache:", err)
	}
	return spinner
}
