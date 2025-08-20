package iam

import (
	aw "github.com/deanishe/awgo"

	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type ServiceAccountSearcher struct{}

func (s *ServiceAccountSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"iam_service_accounts",
		wf,
		cfg,
		q,
		gc.ListIAMServiceAccounts,
		func(wf *aw.Workflow, account gc.IAMServiceAccount) {
			sb := FromGCloudIAMServiceAccount(&account)
			resource.NewItem(wf, cfg, sb, svc.Icon())
		},
	)

	return builder.Build()
}
