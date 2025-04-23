package main

import (
	"embed"
	"flag"
	"log"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/orchestrator"
)

var (
	wf         *aw.Workflow
	query      string
	forceFetch bool

	//go:embed services.yml
	servicesFs embed.FS
)

func init() {
	wf = aw.New()

	flag.StringVar(&query, "query", "", "Query to use")
	flag.BoolVar(&forceFetch, "fetch", false, "Force fetch bypassing cache")
	flag.Parse()
}

func main() {
	orch := orchestrator.NewOrchestrator(
		&orchestrator.HomeHandler{},
		&orchestrator.ServiceHandler{},
		&orchestrator.SubServiceHandler{},
		&orchestrator.FallbackHandler{},
		&orchestrator.ErrorHandler{},
		servicesFs,
	)

	query = strings.TrimSpace(query)
	log.Println("LOG: raw query:", query)

	wf.Run(func() {
		orch.Run(wf, query)
	})
}
