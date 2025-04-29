package artifactregistry

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type RepositorySearcher struct{}

func (s *RepositorySearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, pq *parser.Result) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"artifact_registry_repositories",
		wf,
		config,
		pq,
		s.fetch,
		func(wf *aw.Workflow, entity gc.ArtifactRepository) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *RepositorySearcher) fetch(config *gc.Config) ([]gc.ArtifactRepository, error) {
	return gc.ListArtifactRepositories(config)
}

func (s *RepositorySearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.ArtifactRepository) {
	repo := RepositoryFromGCloud(&entity)
	wf.NewItem(repo.Title()).
		Subtitle(repo.Subtitle()).
		Arg(repo.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
