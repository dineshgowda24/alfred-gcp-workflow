package main

import (
	"embed"
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/cmd"
)

var (
	Version = "dev"
	//go:embed services.yml
	servicesFs embed.FS
)

func init() {
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[Version: " + Version + "] ")
}

func main() {
	wf := cmd.NewWorkflow()
	logDirs(wf)
	cmd.NewRunner(wf, servicesFs).
		RewireMagicQuery().
		ParseFlags().
		Run()
}

func logDirs(wf *aw.Workflow) {
	log.Println("root dir:", wf.Dir())
	log.Println("data dir:", wf.DataDir())
	log.Println("cache dir:", wf.CacheDir())
}
