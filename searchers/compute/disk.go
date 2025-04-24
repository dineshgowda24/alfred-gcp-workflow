package compute

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type DiskSearcher struct{}

func (s *DiskSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.RenderRequest[gcloud.ComputeDisk]{
		Key:    "compute_disks",
		Wf:     wf,
		Config: config,
		Args:   &args,
		Fetch:  s.fetch,
		Render: func(wf *aw.Workflow, entity gcloud.ComputeDisk) {
			s.render(wf, svc, config, entity)
		},
	})
}

func (s *DiskSearcher) fetch(config *gcloud.Config) ([]gcloud.ComputeDisk, error) {
	return gcloud.ListComputeDisks(config)
}

func (s *DiskSearcher) render(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, entity gcloud.ComputeDisk) {
	disk := FromGCloudComputeDisk(&entity)
	wf.NewItem(disk.Title()).
		Subtitle(disk.Subtitle()).
		Arg(disk.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
