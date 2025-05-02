package storage

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type BucketSearcher struct{}

func (s *BucketSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"storage_buckets",
		wf,
		cfg,
		q,
		gc.ListCloudStorageBuckets,
		func(wf *aw.Workflow, gsb gc.Bucket) {
			sb := FromGCloudStorageBucket(&gsb)
			resource.NewItem(wf, cfg, sb, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
