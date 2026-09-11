package resource

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Displayable interface {
	Title() string
	Subtitle() string
	URL(*gcloud.Config) string
}

func NewItem(wf *aw.Workflow, config *gcloud.Config, item Displayable, icon *aw.Icon) *aw.Item {
	return wf.NewItem(item.Title()).
		Subtitle(item.Subtitle()).
		Arg(gcloud.WithAuthUser(item.URL(config), config)).
		Icon(icon).
		Valid(true)
}
