package pubsub

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type SubscriptionSearcher struct{}

func (s *SubscriptionSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.RenderRequest[gcloud.PubSubSubscription]{
		Key:    "pubsub_subscriptions",
		Wf:     wf,
		Config: config,
		Args:   &args,
		Fetch:  s.fetch,
		Render: func(wf *aw.Workflow, entity gcloud.PubSubSubscription) {
			s.render(wf, svc, config, entity)
		},
	})
}

func (s *SubscriptionSearcher) fetch(config *gcloud.Config) ([]gcloud.PubSubSubscription, error) {
	return gcloud.ListSubscriptions(config)
}

func (s *SubscriptionSearcher) render(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, entity gcloud.PubSubSubscription) {
	sub := FromGCloudSubscription(&entity)
	wf.NewItem(sub.Title()).
		Subtitle(sub.Subtitle()).
		Arg(sub.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
