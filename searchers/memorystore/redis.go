package memorystore

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type RedisInstanceSearcher struct{}

func (s *RedisInstanceSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.RenderRequest[gcloud.RedisInstance]{
		Key:    "memorystore_redis_instances",
		Wf:     wf,
		Config: config,
		Args:   &args,
		Fetch:  s.fetch,
		Render: func(wf *aw.Workflow, entity gcloud.RedisInstance) {
			s.render(wf, svc, config, entity)
		},
	})
}

func (s *RedisInstanceSearcher) fetch(config *gcloud.Config) ([]gcloud.RedisInstance, error) {
	return gcloud.ListRedisInstances(config)
}

func (s *RedisInstanceSearcher) render(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, entity gcloud.RedisInstance) {
	ins := RedisInstanceFromGCloud(&entity)
	wf.NewItem(ins.Name).
		Subtitle(ins.Subtitle()).
		Arg(ins.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
