package pubsub

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type TopicSearcher struct{}

func (s *TopicSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.RenderRequest[gcloud.PubSubTopic]{
		Key:    "pubsub_topics",
		Wf:     wf,
		Config: config,
		Args:   &args,
		Fetch:  s.fetch,
		Render: func(wf *aw.Workflow, entity gcloud.PubSubTopic) {
			s.render(wf, svc, config, entity)
		},
	})
}

func (s *TopicSearcher) fetch(config *gcloud.Config) ([]gcloud.PubSubTopic, error) {
	return gcloud.ListTopics(config)
}

func (s *TopicSearcher) render(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, entity gcloud.PubSubTopic) {
	topic := FromGCloudTopic(&entity)
	wf.NewItem(topic.Title()).
		Subtitle(topic.Subtitle()).
		Arg(topic.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
