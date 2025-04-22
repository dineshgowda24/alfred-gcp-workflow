package compute

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type ImageSearcher struct{}

func (s *ImageSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, filter string) error {
	gimages, err := gcloud.ListComputeImages(config)
	if err != nil {
		return err
	}

	for _, gi := range gimages {
		disk := FromGCloudComputeImage(&gi)
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
