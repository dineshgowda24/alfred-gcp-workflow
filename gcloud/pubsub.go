package gcloud

import (
	"encoding/json"
	"os/exec"
)

func ListTopics(config *Config) ([]PubSubTopic, error) {
	cmd := exec.Command(
		gcloudPath,
		"pubsub", "topics", "list",
		"--format=json",
		"--project="+config.Project,
	)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var topics []PubSubTopic
	if err := json.Unmarshal(output, &topics); err != nil {
		return nil, err
	}
	return topics, nil
}

func ListSubscriptions(config *Config) ([]PubSubSubscription, error) {
	cmd := exec.Command(
		gcloudPath,
		"pubsub", "subscriptions", "list",
		"--format=json",
		"--project="+config.Project,
	)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var subs []PubSubSubscription
	if err := json.Unmarshal(output, &subs); err != nil {
		return nil, err
	}
	return subs, nil
}
