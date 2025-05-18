package main

import (
	"embed"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/cmd"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/config"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/log"
)

var (
	Version = "dev"
	//go:embed services.yml
	servicesFs embed.FS
	//go:embed regions.yml
	regionsFs embed.FS
)

func init() {
	gcloud.InitRegions(regionsFs)
}

func main() {
	wf := cmd.NewWorkflow()

	log.Init(Version, wf)
	config.Init(wf)

	logDirs(wf)
	cmd.NewRunner(wf, servicesFs).
		RewireMagicQuery().
		ParseFlags().
		Run()
}

func logDirs(wf *aw.Workflow) {
	log.Debug("root dir:", wf.Dir())
	log.Debug("data dir:", wf.DataDir())
	log.Debug("cache dir:", wf.CacheDir())
}
