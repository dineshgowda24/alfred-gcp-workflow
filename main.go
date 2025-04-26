package main

import (
	"embed"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/cmd"
)

//go:embed services.yml
var servicesFs embed.FS

func main() {
	wf := aw.New(aw.MagicPrefix(cmd.MagicPrefix))
	cmd.NewRunner(wf, servicesFs).
		RewireMagicQuery().
		ParseFlags().
		Run()
}
