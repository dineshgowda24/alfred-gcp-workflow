package gcloud

// https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics#Topic
type PubSubTopic struct {
	Name string `json:"name"` // Format: projects/{project}/topics/{topic-id}
}

// https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions#Subscription
type PubSubSubscription struct {
	AckDeadlineSeconds       int    `json:"ackDeadlineSeconds"`
	MessageRetentionDuration string `json:"messageRetentionDuration"`
	Name                     string `json:"name"` // Format: projects/{project}/subscriptions/{subscription-id}
	State                    string `json:"state"`
	Topic                    string `json:"topic"` // Format: projects/{project}/topics/{topic-id}
}

func ListTopics(config *Config) ([]PubSubTopic, error) {
	return runGCloudCmd[[]PubSubTopic](
		config, "pubsub", "topics", "list",
		"--format=json(name)",
	)
}

func ListSubscriptions(config *Config) ([]PubSubSubscription, error) {
	return runGCloudCmd[[]PubSubSubscription](
		config, "pubsub", "subscriptions", "list",
		"--format=json(ackDeadlineSeconds,messageRetentionDuration,name,state,topic)",
	)
}
