package pubsub

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

type SubscriptionSearcher struct{}

func (s *SubscriptionSearcher) Search(wf *aw.Workflow, svc *services.Service, config *gc.Config, args workflow.SearchArgs) error {
	return workflow.ResolveAndRender(workflow.NewRenderRequest(
		"pubsub_subscriptions",
		wf,
		config,
		&args,
		s.fetch,
		func(wf *aw.Workflow, entity gc.PubSubSubscription) {
			s.render(wf, svc, config, entity)
		},
	))
}

func (s *SubscriptionSearcher) fetch(config *gc.Config) ([]gc.PubSubSubscription, error) {
	return gc.ListSubscriptions(config)
}

func (s *SubscriptionSearcher) render(wf *aw.Workflow, svc *services.Service, config *gc.Config, entity gc.PubSubSubscription) {
	sub := FromGCloudSubscription(&entity)
	wf.NewItem(sub.Title()).
		Subtitle(sub.Subtitle()).
		Arg(sub.URL(config)).
		Icon(svc.Icon()).
		Valid(true)
}
