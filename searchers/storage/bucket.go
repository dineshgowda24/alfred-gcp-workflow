package storage

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type BucketSearcher struct{}

func (s *BucketSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, args workflow.SearchArgs) error {
	return workflow.LoadFromCache(
		wf,
		config.CacheKey("storage_buckets"),
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gcloud.Bucket) {
			s.render(wf, svc, config, entity)
		},
	)
}

func (s *BucketSearcher) fetch(config *gcloud.Config) ([]gcloud.Bucket, error) {
	return gcloud.ListCloudStorageBuckets(config)
}

func (s *BucketSearcher) render(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, entity gcloud.Bucket) {
	buck := FromGCloudStorageBucket(&entity)
	wf.NewItem(buck.Title()).
		Subtitle(buck.Subtitle()).
		Arg(buck.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
