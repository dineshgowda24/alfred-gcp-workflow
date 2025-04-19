package pubsub

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type SubscriptionSearcher struct{}

func (s *SubscriptionSearcher) Search(wf *aw.Workflow, config *gcloud.Config, filter string) error {
	gsubs, err := gcloud.ListSubscriptions(config)
	if err != nil {
		wf.NewItem("❌ Error loading subscriptions").
			Subtitle(err.Error()).
			Valid(false)
		return nil
	}

	for _, gsub := range gsubs {
		sub := FromGCloudSubscription(&gsub)

		wf.NewItem(sub.Title()).
			Subtitle(sub.Subtitle()).
			Arg(sub.URL(config)).
			Valid(true)
	}

	if filter != "" {
		wf.Filter(filter)
	}
	return nil
}
