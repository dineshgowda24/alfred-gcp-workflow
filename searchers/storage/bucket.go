package storage

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type BucketSearcher struct{}

func (s *BucketSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"storage_buckets",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.Bucket) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *BucketSearcher) fetch(config *gc.Config) ([]gc.Bucket, error) {
	return gc.ListCloudStorageBuckets(config)
}

func (s *BucketSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.Bucket) {
	buck := FromGCloudStorageBucket(&entity)
	wf.NewItem(buck.Title()).
		Subtitle(buck.Subtitle()).
		Arg(buck.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
