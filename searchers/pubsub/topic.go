package pubsub

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type TopicSearcher struct{}

func (s *TopicSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"pubsub_topics",
		wf,
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gc.PubSubTopic) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *TopicSearcher) fetch(config *gc.Config) ([]gc.PubSubTopic, error) {
	return gc.ListTopics(config)
}

func (s *TopicSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.PubSubTopic) {
	topic := FromGCloudTopic(&entity)
	wf.NewItem(topic.Title()).
		Subtitle(topic.Subtitle()).
		Arg(topic.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
