package orchestrator

import (
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/arg"
)

// Context holds the context for the workflow execution.
type Context struct {
	Args           *arg.SearchArgs
	Workflow       *aw.Workflow
	ParsedQuery    *parser.Result
	ActiveConfig   *gcloud.Config
	Services       []services.Service
	Err            error
	SearchRegistry *searchers.Registry
}

func (ctx *Context) IsHomeQuery() bool {
	return strings.TrimSpace(ctx.Args.Query) == ""
}

func (ctx *Context) IsConfigQuery() bool {
	return ctx.ParsedQuery.IsConfigQuery
}

func (ctx *Context) IsConfigActive() bool {
	return ctx.ParsedQuery.Config != nil
}

func (ctx *Context) IsRegionQuery() bool {
	return ctx.ParsedQuery.IsRegionQuery
}

func (ctx *Context) IsRegionActive() bool {
	return ctx.ParsedQuery.Region != nil
}

func (ctx *Context) IsServiceQuery() bool {
	return ctx.ParsedQuery.HasServiceOnly()
}

func (ctx *Context) IsSubServiceQuery() bool {
	return ctx.ParsedQuery.HasSubService()
}

func (ctx *Context) IsIntentQuery() bool {
	return ctx.ParsedQuery.HasIntent()
}

func (ctx *Context) SendFeedback() {
	if ctx.Err != nil {
		ctx.Workflow.NewItem("Error: " + ctx.Err.Error()).
			Subtitle("Please check the logs for more details").
			Valid(false)
	}
	ctx.Workflow.SendFeedback()
}

func (ctx *Context) GetSearcher(parent, child *services.Service) searchers.Searcher {
	if ctx.SearchRegistry == nil {
		ctx.SearchRegistry = searchers.GetDefaultRegistry()
	}
	return ctx.SearchRegistry.Get(parent, child)
}
