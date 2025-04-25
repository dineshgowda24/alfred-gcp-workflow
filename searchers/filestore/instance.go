package filestore

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type InstanceSearcher struct{}

func (s *InstanceSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"filestore_instances",
		wf,
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gc.FilestoreInstance) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *InstanceSearcher) fetch(config *gc.Config) ([]gc.FilestoreInstance, error) {
	return gc.ListFilestoreInstances(config)
}

func (s *InstanceSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.FilestoreInstance) {
	ins := InstanceFromGCloud(&entity)
	wf.NewItem(ins.Id).
		Subtitle(ins.Subtitle()).
		Arg(ins.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
