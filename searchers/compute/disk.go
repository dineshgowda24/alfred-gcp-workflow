package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type DiskSearcher struct{}

func (s *DiskSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"compute_disks",
		wf,
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gc.ComputeDisk) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *DiskSearcher) fetch(config *gc.Config) ([]gc.ComputeDisk, error) {
	return gc.ListComputeDisks(config)
}

func (s *DiskSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.ComputeDisk) {
	disk := FromGCloudComputeDisk(&entity)
	wf.NewItem(disk.Title()).
		Subtitle(disk.Subtitle()).
		Arg(disk.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
