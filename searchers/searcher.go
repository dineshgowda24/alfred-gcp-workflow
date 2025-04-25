package searchers

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/compute"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/filestore"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/k8s"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/memorystore"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/netservices"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/pubsub"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/sql"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/storage"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/vpc"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type Searcher interface {
	Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, pq *parser.Result) error
}

type Registry struct {
	lookup map[string]Searcher
}

func (r *Registry) Get(parent, child *services.Service) Searcher {
	searcher, ok := r.lookup[r.key(parent, child)]
	if !ok {
		return nil
	}

	return searcher
}

func (r *Registry) key(parent, child *services.Service) string {
	return parent.ID + "/" + child.ID
}

func (r *Registry) Exists(parent, child *services.Service) bool {
	_, ok := r.lookup[r.key(parent, child)]
	return ok
}

func GetDefaultRegistry() *Registry {
	return &Registry{
		lookup: map[string]Searcher{
			"sql/instances":              &sql.InstanceSearcher{},
			"memorystore/redis":          &memorystore.RedisInstanceSearcher{},
			"pubsub/topics":              &pubsub.TopicSearcher{},
			"pubsub/subscriptions":       &pubsub.SubscriptionSearcher{},
			"storage/buckets":            &storage.BucketSearcher{},
			"compute/instances":          &compute.InstanceSearcher{},
			"compute/disks":              &compute.DiskSearcher{},
			"compute/images":             &compute.ImageSearcher{},
			"compute/instance_templates": &compute.InstanceTmplSearcher{},
			"gke/clusters":               &k8s.ClusterSearcher{},
			"filestore/instances":        &filestore.InstanceSearcher{},
			"net_svc/cloud_dns":          &netservices.DNSZoneSearcher{},
			"vpc/networks":               &vpc.NetworkSearcher{},
			"vpc/routes":                 &vpc.RouteSearcher{},
		},
	}
}
