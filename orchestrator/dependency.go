package orchestrator

import (
	"errors"
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/env"
)

type DependencyHandler struct{}

func (h *DependencyHandler) Handle(ctx *Context) error {
	wf := ctx.Workflow

	if env.GCloudCliPath() == "" {
		wf.NewItem("gcloud path not set").
			Subtitle(fmt.Sprintf("Please set %s in alfred", env.EnvGCloudCliPath)).
			Icon(aw.IconError).
			Valid(false)
		wf.SendFeedback()
		return errors.New("gcloud path not set")
	}

	return nil
}
