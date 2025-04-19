package sql

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type InstanceSearcher struct{}

func (s *InstanceSearcher) Search(wf *aw.Workflow, config *gcloud.Config, filter string) error {
	ginstances, err := gcloud.ListSQLInstances(config)
	if err != nil {
		return err
	}

	for _, ginst := range ginstances {
		inst := FromGCloudSQLInstance(&ginst)
		wf.NewItem(inst.Title()).
			Subtitle(inst.Subtitle()).
			Arg(inst.URL(config)).
			Valid(true)
	}

	if filter != "" {
		wf.Filter(filter)
	}

	return nil
}
