package compute

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type InstanceSearcher struct{}

func (s *InstanceSearcher) Search(
	wf *aw.Workflow,
	svc *services.Service,
	config *gcloud.Config,
	args workflow.SearchArgs,
) error {
	return workflow.LoadFromCache(
		wf,
		config.CacheKey("compute_instances"),
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gcloud.ComputeInstance) {
			s.render(wf, svc, config, entity)
		},
	)
}

func (s *InstanceSearcher) fetch(config *gcloud.Config) ([]gcloud.ComputeInstance, error) {
	return gcloud.ListComputeInstances(config)
}

func (s *InstanceSearcher) render(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, entity gcloud.ComputeInstance) {
	inst := FromGCloudComputeInstance(&entity)
	wf.NewItem(inst.Title()).
		Subtitle(inst.Subtitle()).
		Arg(inst.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
