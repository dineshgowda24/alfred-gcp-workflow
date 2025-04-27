package orchestrator

import (
	"embed"
	"errors"
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/arg"
)

var ErrNoActiveConfig = errors.New("no active gcloud config found")

type Handler interface {
	Handle(ctx *Context) error
}

type Orchestrator struct {
	ctx               *Context
	preflightHandler  Handler
	configHandler     Handler
	homeHandler       Handler
	serviceHandler    Handler
	subServiceHandler Handler
	fallback          Handler
	errorHandler      Handler
	servicesFS        embed.FS
}

func DefaultOrchestrator(servicesFS embed.FS) *Orchestrator {
	return NewOrchestrator(
		&PreFlightCheckHandler{},
		&ConfigHandler{},
		&HomeHandler{},
		&ServiceHandler{},
		&SubServiceHandler{},
		&FallbackHandler{},
		&ErrorHandler{},
		servicesFS,
	)
}

func NewOrchestrator(preflight, config, home, service, subService, fallback, err Handler, servicesFS embed.FS) *Orchestrator {
	return &Orchestrator{
		preflightHandler:  preflight,
		configHandler:     config,
		homeHandler:       home,
		serviceHandler:    service,
		subServiceHandler: subService,
		fallback:          fallback,
		errorHandler:      err,
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

	log.Println("LOG: IsConfigQuery:", o.ctx.IsConfigQuery())
	log.Println("LOG: IsConfigActive:", o.ctx.IsConfigActive())

	switch {
	case (o.ctx.IsConfigQuery() && !o.ctx.IsConfigActive()):
		o.ctx.Err = o.configHandler.Handle(o.ctx)
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
	servicesList, err := services.Load(o.servicesFS)
	if err != nil {
		o.ctx.Err = err
		return
	}

	o.ctx.Services = servicesList
	o.ctx.SearchRegistry = searchers.GetDefaultRegistry()
	o.ctx.ParsedQuery = parser.Parse(args, servicesList)

	if o.ctx.ParsedQuery.Config != nil {
		o.ctx.ActiveConfig = o.ctx.ParsedQuery.Config
	} else {
		config := gcloud.GetActiveConfig()
		if config == nil {
			o.ctx.Err = ErrNoActiveConfig
			return
		}

		o.ctx.ActiveConfig = config
	}
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
