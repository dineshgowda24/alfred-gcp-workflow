package resource

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/env"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/log"
)

const (
	backgroundJob = "alfred-gcp-workflow-background-job"
	// defaultCacheTTL is the default time to live for the cache.
	defaultCacheTTL = 7 * 24 * time.Hour
	// errCacheTTL is the time to live for the error cache.
	errCacheTTL = 5 * time.Second
)

var (
	cacheTTL              = env.CacheTTLDuration(defaultCacheTTL)
	ErrCacheExpired       = errors.New("cache expired")
	ErrRegionNotSupported = errors.New("region not supported")
)

type (
	// Fetcher is a function that fetches resources from gcloud.
	// It takes a gcloud config and returns a slice of resources or an error.
	Fetcher[T any] func(*gc.Config) ([]T, error)
	// Renderer is a function that renders a resource to the Alfred workflow.
	// It takes a workflow and a resource and adds it to the workflow.
	Renderer[T any] func(*aw.Workflow, T)
)

// RegionSupporter defines an optional interface for resources that can validate region support.
// It's used to proactively check if a region (or location) is valid for a given gcloud service.
//
// This is useful when users override the region via Script Filter arguments.
// In some cases, using an unsupported region causes gcloud commands to fail, resulting in poor UX.
// If a resource implements RegionSupporter, the workflow will check region validity before calling gcloud,
// allowing the user to be informed immediately.
type RegionSupporter interface {
	IsRegionSupported(config *gc.Config) bool
}

type UXMsg struct {
	Error    string `json:"error,omitempty"`
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
}

// Builder is a struct that builds a gcloud resource item for Alfred.
type Builder[T any] struct {
	key     string
	wf      *aw.Workflow
	cfg     *gc.Config
	query   *parser.Result
	fetch   Fetcher[T]
	render  Renderer[T]
	spinner *spinner
}

func NewBuilder[T any](
	key string, wf *aw.Workflow, cfg *gc.Config, q *parser.Result, fetch Fetcher[T], render Renderer[T],
) *Builder[T] {
	return &Builder[T]{
		key:     key,
		wf:      wf,
		cfg:     cfg,
		query:   q,
		fetch:   fetch,
		render:  render,
		spinner: defaultSpinner(wf),
	}
}

// Build fetches the resources from gcloud or loads them from the cache and renders them to the alfred.
// If the cache exists and is not expired, it loads the resources from the cache and renders them.
// If the cache is expired or not found, it starts a background job to fetch the resources. The background job populates the cache.
// The workflow is instructed to rerun after a short delay to allow the background job to complete.
func (b *Builder[T]) Build() error {
	if b.query.SearchArgs.RebuildCache {
		return b.store()
	}

	if b.tryCache() {
		return nil
	}

	b.loading()
	b.runBackgroundJob()
	b.wf.Rerun(0.2)
	return nil
}

func (b *Builder[T]) store() error {
	if !b.validateRegion() {
		return ErrRegionNotSupported
	}

	data, err := b.fetch(b.cfg)
	if err != nil {
		msg := UXMsg{
			Error: err.Error(),
		}
		b.wf.Cache.StoreJSON(b.errKey(), &msg)
		return err
	}
	return b.wf.Cache.StoreJSON(b.cacheKey(), data)
}

func (b *Builder[T]) validateRegion() bool {
	var resource T
	supporter, ok := any(resource).(RegionSupporter)
	if ok && supporter != nil && !supporter.IsRegionSupported(b.cfg) {
		msg := UXMsg{
			Title:    "🔔 Selected region not supported for this resource",
			Subtitle: "May be try another region? 🤷",
		}

		b.wf.Cache.StoreJSON(b.errKey(), &msg)
		return false
	}

	return true
}

func (b *Builder[T]) errKey() string {
	return b.cfg.CacheKey(b.key) + "_gcloud_error"
}

func (b *Builder[T]) cacheKey() string {
	return b.cfg.CacheKey(b.key)
}

func (b *Builder[T]) tryCache() bool {
	if b.showCachedErr() {
		return true
	}

	resources, err := b.load()
	if err != nil {
		log.Error("failed to load cache:", err)
		return false
	}

	b.renderAll(resources...)
	return true
}

func (b *Builder[T]) showCachedErr() bool {
	var msg UXMsg
	if !b.wf.Cache.Expired(b.errKey(), errCacheTTL) {
		err := b.wf.Cache.LoadJSON(b.errKey(), &msg)
		if err != nil {
			log.Error("failed to load cached error:", err)

			if err := b.delCacheKey(); err != nil {
				log.Error("failed to delete cache file:", err)
			}

			msg = UXMsg{
				Error:    "🔔 Failed loading cached error",
				Subtitle: err.Error(),
			}
		}

		title := msg.Title
		icon := aw.IconInfo
		subtitle := msg.Subtitle
		if msg.Error != "" {
			title = "🔔 Error fetching resources"
			subtitle = "<CMD> + l to view error details"
			icon = aw.IconError
		}

		b.wf.NewItem(title).
			Subtitle(subtitle).
			Icon(icon).
			Arg("").
			Valid(false).
			Largetype(msg.Error).
			Autocomplete("")

		if msg.Error != "" {
			errorFilePath := filepath.Join(b.wf.CacheDir(), b.errKey())
			b.wf.NewFileItem("📂 View full error file").
				Subtitle("<CTRL> to open file options").
				Icon(aw.IconInfo).
				IsFile(true).
				Arg(errorFilePath).
				Valid(true)

		}
		b.wf.SendFeedback()
		return true
	}
	return false
}

func (b *Builder[T]) load() ([]T, error) {
	key := b.cacheKey()
	if !b.wf.Cache.Expired(key, cacheTTL) {
		log.Debug("cache found")
		return b.loadJSONCache()
	}

	// Cache expired, but an expired cache file exists.
	// Run the background job to refresh the cache and return the expired cache.
	// This allows the user to see the stale data while the background job fetches fresh data.
	if b.wf.Cache.Exists(key) {
		log.Info("cache expired, refreshing in background")
		b.runBackgroundJob()
		log.Debug("returning expired cache")
		return b.loadJSONCache()
	}

	log.Debug("cache expired or not found")
	return nil, ErrCacheExpired
}

func (b *Builder[T]) loadJSONCache() ([]T, error) {
	v := make([]T, 0)
	if err := b.wf.Cache.LoadJSON(b.cacheKey(), &v); err != nil {
		log.Error("failed to load cache:", err)
		if b.delCacheKey() != nil {
			return nil, err
		}
		return nil, err
	}

	return v, nil
}

func (b *Builder[T]) delCacheKey() error {
	if err := os.Remove(filepath.Clean(filepath.Join(b.wf.CacheDir(), b.cacheKey()))); err != nil {
		if !os.IsNotExist(err) {
			log.Error("failed to delete cache file:", err)
			return err
		}
	}
	return nil
}

func (b *Builder[T]) renderAll(resources ...T) {
	if len(resources) == 0 {
		b.wf.NewItem("No resources found").
			Subtitle("Try a different query").
			Icon(aw.IconInfo).
			Valid(false)
	}

	for _, resource := range resources {
		b.render(b.wf, resource)
	}
}

func (b *Builder[T]) loading() {
	b.wf.NewItem(b.spinner.NextFrame() + " Fetching resources").
		Subtitle("Hang tight 🙏").
		Icon(aw.IconInfo).
		Valid(false)
}

func (b *Builder[T]) runBackgroundJob() {
	if b.wf.IsRunning(backgroundJob) {
		return
	}
	cmd := exec.Command(os.Args[0], "--query="+b.query.SearchArgs.Query, "--rebuild-cache")
	if err := b.wf.RunInBackground(backgroundJob, cmd); err != nil {
		panic(err)
	}
}
