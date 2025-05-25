package searchers

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/artifactregistry"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/cloudrun"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/cloudtask"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/compute"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/filestore"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/k8s"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/memorystore"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/netconnectivity"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/netservices"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/pubsub"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/sql"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/storage"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/vpc"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

// Searcher is an interface for searching GCP resources.
type Searcher interface {
	Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, pq *parser.Result) error
}

// Registry contains the map of all the searchers.
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

// GetDefaultRegistry returns a default registry with all the searchers registered.
func GetDefaultRegistry() *Registry {
	return &Registry{
		lookup: map[string]Searcher{
			"artifactregistry/repositories": &artifactregistry.RepositorySearcher{},
			"cloudrun/functions":            &cloudrun.FunctionSearcher{},
			"cloudrun/services":             &cloudrun.ServiceSearcher{},
			"cloudsql/instances":            &sql.InstanceSearcher{},
			"cloudtasks/queues":             &cloudtask.QueueSearcher{},
			"compute/disks":                 &compute.DiskSearcher{},
			"compute/images":                &compute.ImageSearcher{},
			"compute/instances":             &compute.InstanceSearcher{},
			"compute/instancetemplates":     &compute.InstanceTmplSearcher{},
			"compute/machineimages":         &compute.MachineImageSearcher{},
			"compute/snapshots":             &compute.SnapshotSearcher{},
			"filestore/instances":           &filestore.InstanceSearcher{},
			"gke/clusters":                  &k8s.ClusterSearcher{},
			"memorystore/redis":             &memorystore.RedisInstanceSearcher{},
			"netconnectivity/cloudrouter":   &netconnectivity.CloudRouterSearcher{},
			"netconnectivity/vpngateway":    &netconnectivity.VPNGatewaySearcher{},
			"netconnectivity/vpntunnel":     &netconnectivity.VPNTunnelSearcher{},
			"netservices/dns":               &netservices.DNSZoneSearcher{},
			"pubsub/subscriptions":          &pubsub.SubscriptionSearcher{},
			"pubsub/topics":                 &pubsub.TopicSearcher{},
			"storage/buckets":               &storage.BucketSearcher{},
			"vpc/networks":                  &vpc.NetworkSearcher{},
			"vpc/routes":                    &vpc.RouteSearcher{},
		},
	}
}
