package cmd

import (
	"embed"
	"flag"

	aw "github.com/deanishe/awgo"
	ors "github.com/dineshgowda24/alfred-gcp-workflow/orchestrator"
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

func buildOrch(servicesFs embed.FS) *ors.Orchestrator {
	return ors.NewOrchestrator(
		&ors.PreFlightCheckHandler{},
		&ors.HomeHandler{},
		&ors.ServiceHandler{},
		&ors.SubServiceHandler{},
		&ors.FallbackHandler{},
		&ors.ErrorHandler{},
		servicesFs,
	)
}
