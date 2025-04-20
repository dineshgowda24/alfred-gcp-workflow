package storage

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type BucketSearcher struct{}

func (s *BucketSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, filter string) error {
	gbuckets, err := gcloud.ListCloudStorageBuckets(config)
	if err != nil {
		return err
	}

	for _, ginst := range gbuckets {
		buck := FromGCloudStorageBucket(&ginst)
		wf.NewItem(buck.Title()).
			Subtitle(buck.Subtitle()).
			Arg(buck.URL(config)).
			Icon(svc.Icon()).
			Valid(true)
	}

	if filter != "" {
		wf.Filter(filter)
	}

	return nil
}
