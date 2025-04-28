package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type SnapshotSearcher struct{}

func (s *SnapshotSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"compute_snapshots",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.ComputeSnapshot) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *SnapshotSearcher) fetch(config *gc.Config) ([]gc.ComputeSnapshot, error) {
	return gc.ListComputeSnapshots(config)
}

func (s *SnapshotSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.ComputeSnapshot) {
	snapshot := FromGCloudComputeSnapshot(&entity)
	wf.NewItem(snapshot.Title()).
		Subtitle(snapshot.Subtitle()).
		Arg(snapshot.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
