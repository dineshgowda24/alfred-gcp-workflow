package compute

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type DiskSearcher struct{}

func (s *DiskSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, filter string) error {
	gdisks, err := gcloud.ListComputeDisks(config)
	if err != nil {
		return err
	}

	for _, gd := range gdisks {
		disk := FromGCloudComputeDisk(&gd)
		wf.NewItem(disk.Title()).
			Subtitle(disk.Subtitle()).
			Arg(disk.URL(config)).
			Icon(svc.Icon()).
			Valid(true)
	}

	if filter != "" {
		wf.Filter(filter)
	}

	return nil
}
