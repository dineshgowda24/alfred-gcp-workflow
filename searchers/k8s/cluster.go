package k8s

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type ClusterSearcher struct{}

func (s *ClusterSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"k8s_clusters",
		wf,
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gc.K8sCluster) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *ClusterSearcher) fetch(config *gc.Config) ([]gc.K8sCluster, error) {
	return gc.ListK8sClusters(config)
}

func (s *ClusterSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.K8sCluster) {
	cluster := ClusterFromGCloud(&entity)
	wf.NewItem(cluster.Title()).
		Subtitle(cluster.Subtitle()).
		Arg(cluster.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
