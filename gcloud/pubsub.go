package gcloud

type PubSubTopic struct {
	Name string `json:"name"` // Format: projects/{project}/topics/{topic-id}
}

type PubSubSubscription struct {
	AckDeadlineSeconds       int    `json:"ackDeadlineSeconds"`
	MessageRetentionDuration string `json:"messageRetentionDuration"`
	Name                     string `json:"name"` // Format: projects/{project}/subscriptions/{subscription-id}
	State                    string `json:"state"`
	Topic                    string `json:"topic"` // Format: projects/{project}/topics/{topic-id}
}

func ListTopics(config *Config) ([]PubSubTopic, error) {
	return runGCloudCmd[[]PubSubTopic](config, "pubsub", "topics", "list")
}

func ListSubscriptions(config *Config) ([]PubSubSubscription, error) {
	return runGCloudCmd[[]PubSubSubscription](config, "pubsub", "subscriptions", "list")
}
