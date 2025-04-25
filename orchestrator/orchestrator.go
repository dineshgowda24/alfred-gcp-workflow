package orchestrator

import (
	"embed"
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/arg"
)

type Handler interface {
	Handle(ctx *Context) error
}

type Orchestrator struct {
	ctx               *Context
	preflightHandler  Handler
	homeHandler       Handler
	serviceHandler    Handler
	subServiceHandler Handler
	fallback          Handler
	errorHandler      Handler
	servicesFS        embed.FS
}

func NewOrchestrator(preflight, home, service, subService, fallback, error Handler, servicesFS embed.FS) *Orchestrator {
	return &Orchestrator{
		preflightHandler:  preflight,
		homeHandler:       home,
		serviceHandler:    service,
		subServiceHandler: subService,
		fallback:          fallback,
		errorHandler:      error,
		servicesFS:        servicesFS,
	}
}

func (o *Orchestrator) Run(wf *aw.Workflow, args *arg.SearchArgs) {
	log.Println("LOG: orchestrator run with query:", args.Query)

	o.buildCtx(wf, args)
	if o.ctx.Err != nil {
		o.handleErr()
		return
	}
	log.Println("LOG: build context complete", o.ctx.ActiveConfig, o.ctx.ParsedQuery)

	if !o.preflightCheck() {
		return
	}

	switch {
	case o.ctx.IsHomeQuery():
		o.ctx.Err = o.homeHandler.Handle(o.ctx)
	case o.ctx.IsSubServiceQuery():
		o.ctx.Err = o.subServiceHandler.Handle(o.ctx)
	case o.ctx.IsServiceQuery():
		o.ctx.Err = o.serviceHandler.Handle(o.ctx)
	default:
		o.ctx.Err = o.fallback.Handle(o.ctx)
	}

	o.handleErr()
}

func (o *Orchestrator) buildCtx(wf *aw.Workflow, args *arg.SearchArgs) {
	o.ctx = &Context{Workflow: wf, Args: args}
	config := getActiveConfig(o.ctx, o.ctx.IsHomeQuery())
	if config == nil {
		o.ctx.Err = ErrNoActiveConfig
		return
	}

	o.ctx.ActiveConfig = config
	servicesList, err := services.Load(o.servicesFS)
	if err != nil {
		o.ctx.Err = err
		return
	}

	o.ctx.ParsedQuery = parser.Parse(args, servicesList)
	o.ctx.Services = servicesList
	o.ctx.SearchRegistry = searchers.GetDefaultRegistry()
}

func (o *Orchestrator) handleErr() {
	if o.ctx.Err != nil {
		o.errorHandler.Handle(o.ctx)
	}
}

func (o *Orchestrator) preflightCheck() bool {
	err := o.preflightHandler.Handle(o.ctx)
	if err != nil {
		log.Println("LOG: preflight check failed:", err)
		return false
	}
	return true
}
