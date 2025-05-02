package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type SnapshotSearcher struct{}

func (s *SnapshotSearcher) Search(wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result) error {
	builder := resource.NewBuilder(
		"compute_snapshots",
		wf,
		cfg,
		q,
		gc.ListComputeSnapshots,
		func(wf *aw.Workflow, gcs gc.ComputeSnapshot) {
			cs := FromGCloudComputeSnapshot(&gcs)
			resource.NewItem(wf, cfg, cs, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
