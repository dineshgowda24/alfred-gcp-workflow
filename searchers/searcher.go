package searchers

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/compute"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/memorystore"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/pubsub"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/sql"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/storage"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type Searcher interface {
	Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, filter string) error
}

type Registery struct {
	lookup map[string]Searcher
}

func (r *Registery) Get(parent, child *services.Service) Searcher {
	searcher, ok := r.lookup[r.key(parent, child)]
	if !ok {
		return nil
	}

	return searcher
}

func (r *Registery) key(parent, child *services.Service) string {
	return parent.ID + "/" + child.ID
}

func GetDefaultRegistry() *Registery {
	return &Registery{
		lookup: map[string]Searcher{
			"sql/instances":              &sql.InstanceSearcher{},
			"memorystore/redis":          &memorystore.RedisInstanceSearcher{},
			"pubsub/topics":              &pubsub.TopicSearcher{},
			"pubsub/subscriptions":       &pubsub.SubscriptionSearcher{},
			"storage/buckets":            &storage.BucketSearcher{},
			"compute/instances":          &compute.InstanceSearcher{},
			"compute/disks":              &compute.DiskSearcher{},
			"compute/images":             &compute.ImageSearcher{},
			"compute/instance_templates": &compute.InstanceTemplateSearcher{},
		},
	}
}
