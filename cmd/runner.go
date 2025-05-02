package cmd

import (
	"embed"
	"flag"
	"os"
	"strings"

	aw "github.com/deanishe/awgo"
	ors "github.com/dineshgowda24/alfred-gcp-workflow/orchestrator"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/arg"
)

func NewWorkflow() *aw.Workflow {
	return aw.New(aw.MagicPrefix(ors.MagicPrefix))
}

type Runner struct {
	wf         *aw.Workflow
	servicesFS embed.FS
	args       *arg.SearchArgs
}

var (
	query        string
	rebuildCache bool
)

func NewRunner(wf *aw.Workflow, servicesFS embed.FS) *Runner {
	return &Runner{
		wf:         wf,
		servicesFS: servicesFS,
	}
}

func (r *Runner) RewireMagicQuery() *Runner {
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "--query=") {
			query := strings.TrimPrefix(arg, "--query=")
			if strings.HasPrefix(query, ors.MagicPrefix) {
				os.Args[i] = query
			}
		}
	}
	return r
}

func (r *Runner) ParseFlags() *Runner {
	flag.StringVar(&query, "query", "", "User query input")
	flag.BoolVar(&rebuildCache, "rebuild-cache", false, "Force cache rebuild")
	flag.Parse()

	r.args = &arg.SearchArgs{
		Query:        query,
		RebuildCache: rebuildCache,
	}
	return r
}

func (r *Runner) Run() {
	r.interceptMagicArgs()
	r.wf.Run(func() {
		ors.DefaultOrchestrator(r.servicesFS).Run(r.wf, r.args)
	})
}

func (r *Runner) interceptMagicArgs() {
	r.wf.Args()
}
