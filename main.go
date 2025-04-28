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
	wf := aw.New(aw.MagicPrefix(cmd.MagicPrefix))
	cmd.NewRunner(wf, servicesFs).
		RewireMagicQuery().
		ParseFlags().
		Run()
}
