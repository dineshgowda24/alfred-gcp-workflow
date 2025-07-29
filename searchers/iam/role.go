package iam

import (
	aw "github.com/deanishe/awgo"

	gc "github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/resource"
)

type RoleSearcher struct{}

func (s *RoleSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"iam_roles",
		wf,
		cfg,
		q,
		gc.ListIAMRoles,
		func(wf *aw.Workflow, role gc.IAMRole) {
			sb := FromGCloudIAMRoles(&role)
			resource.NewItem(wf, cfg, sb, svc.Icon())
		},
	)

	return builder.Build()
}
