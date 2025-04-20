package memorystore

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type RedisInstanceSearcher struct{}

func (s *RedisInstanceSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, filter string) error {
	gcloudInstances, err := gcloud.ListRedisInstances(config)
	if err != nil {
		wf.NewItem("Error listing Redis instances").
			Subtitle(err.Error()).
			Valid(false)
		return nil
	}

	for _, gi := range gcloudInstances {
		ins := RedisInstanceFromGCloud(&gi)
		wf.NewItem(ins.Name).
			Subtitle(ins.Subtitle()).
			Arg(ins.URL(config)).
			Icon(svc.Icon()).
			Valid(true)
	}

	if filter != "" {
		wf.Filter(filter)
	}

	return nil
}
