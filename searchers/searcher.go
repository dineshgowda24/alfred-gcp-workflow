package searchers

import (
	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/memorystore"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/pubsub"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/sql"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers/storage"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type Searcher interface {
	Search(wf *aw.Workflow, svc *services.Service, config *gcloud.Config, filter string) error
}

var SubserviceSearchers = map[string]Searcher{
	"sql_instances":        &sql.InstanceSearcher{},
	"memorystore_redis":    &memorystore.RedisInstanceSearcher{},
	"pubsub_topics":        &pubsub.TopicSearcher{},
	"pubsub_subscriptions": &pubsub.SubscriptionSearcher{},
	"storage_buckets":      &storage.BucketSearcher{},
}
