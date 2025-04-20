package pubsub

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type TopicSearcher struct{}

func (s *TopicSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, filter string) error {
	gtopics, err := gcloud.ListTopics(config)
	if err != nil {
		return err
	}

	for _, gtopic := range gtopics {
		topic := FromGCloudTopic(&gtopic)

		wf.NewItem(topic.Title()).
			Subtitle(topic.Subtitle()).
			Arg(topic.URL(config)).
			Icon(svc.Icon()).
			Valid(true)
	}

	if filter != "" {
		wf.Filter(filter)
	}
	return nil
}
