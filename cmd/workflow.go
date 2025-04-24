package cmd

import (
	"embed"
	"flag"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/orchestrator"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow"
)

var (
	query        string
	rebuildCache bool
)

func ParseArgs() *workflow.SearchArgs {
	flag.StringVar(&query, "query", "", "Query to use")
	flag.BoolVar(&rebuildCache, "rebuild-cache", false, "Rebuild the disk cache")
	flag.Parse()

	return &workflow.SearchArgs{
		Query:        query,
		RebuildCache: rebuildCache,
	}
}

func RunWorkflow(wf *aw.Workflow, servicesFs embed.FS, args *workflow.SearchArgs) {
	orch := buildOrch(servicesFs)
	wf.Run(func() {
		orch.Run(wf, args)
	})
}

func buildOrch(servicesFs embed.FS) *orchestrator.Orchestrator {
	return orchestrator.NewOrchestrator(
		&orchestrator.HomeHandler{},
		&orchestrator.ServiceHandler{},
		&orchestrator.SubServiceHandler{},
		&orchestrator.FallbackHandler{},
		&orchestrator.ErrorHandler{},
		servicesFs,
	)
}
