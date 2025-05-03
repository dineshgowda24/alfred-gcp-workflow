package gcloud

type PubSubTopic struct {
	Name string `json:"name"`
}

type PubSubSubscription struct {
	AckDeadlineSeconds       int    `json:"ackDeadlineSeconds"`
	MessageRetentionDuration string `json:"messageRetentionDuration"`
	Name                     string `json:"name"`
	State                    string `json:"state"`
	Topic                    string `json:"topic"`
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
