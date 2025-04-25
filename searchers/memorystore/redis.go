package memorystore

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type RedisInstanceSearcher struct{}

func (s *RedisInstanceSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"memorystore_redis_instances",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.RedisInstance) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *RedisInstanceSearcher) fetch(config *gc.Config) ([]gc.RedisInstance, error) {
	return gc.ListRedisInstances(config)
}

func (s *RedisInstanceSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.RedisInstance) {
	ins := RedisInstanceFromGCloud(&entity)
	wf.NewItem(ins.Name).
		Subtitle(ins.Subtitle()).
		Arg(ins.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
