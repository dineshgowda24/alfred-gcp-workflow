package pubsub

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type SubscriptionSearcher struct{}

func (s *SubscriptionSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"pubsub_subscriptions",
		wf,
		cfg,
		q,
		gc.ListSubscriptions,
		func(wf *aw.Workflow, gps gc.PubSubSubscription) {
			ps := FromGCloudSubscription(&gps)
			resource.NewItem(wf, cfg, ps, svc.Icon(wf.Dir()))
		},
	)

	return builder.Build()
}
