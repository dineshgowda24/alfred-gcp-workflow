package main

import (
	"embed"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/cmd"
)

//go:embed services.yml
var servicesFs embed.FS

func main() {
	cmd.RunWorkflow(aw.New(), servicesFs, cmd.ParseArgs())
}
