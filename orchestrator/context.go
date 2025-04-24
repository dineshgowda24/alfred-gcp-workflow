package orchestrator

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

// Context holds the context for the workflow execution.
type Context struct {
	Workflow       *aw.Workflow
	RawQuery       string
	ParsedQuery    *parser.Result
	ActiveConfig   *gcloud.Config
	Services       []services.Service
	Err            error
	SearchRegistry *searchers.Registery
}

func (ctx *Context) IsHomeQuery() bool {
	return ctx.ParsedQuery.IsEmptyQuery()
}

func (ctx *Context) IsServiceQuery() bool {
	return ctx.ParsedQuery.HasServiceOnly()
}

func (ctx *Context) IsSubServiceQuery() bool {
	return ctx.ParsedQuery.HasSubService()
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
