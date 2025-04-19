package pubsub

import (
	"fmt"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Topic struct {
	Name string
}

func FromGCloudTopic(topic *gcloud.PubSubTopic) Topic {
	return Topic{
		Name: topic.Name,
	}
}

func (t Topic) Title() string {
	parts := strings.Split(t.Name, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return t.Name
}

func (t Topic) Subtitle() string {
	return "Pub/Sub Topic"
}

func (t Topic) URL(config *gcloud.Config) string {
	return "https://console.cloud.google.com/cloudpubsub/topic/detail/" + t.Title() + "?project=" + config.Project
}

type Subscription struct {
	Name                     string
	State                    string
	Topic                    string
	AckDeadlineSeconds       int
	MessageRetentionDuration string
}

func FromGCloudSubscription(sub *gcloud.PubSubSubscription) Subscription {
	return Subscription{
		Name:                     sub.Name,
		State:                    sub.State,
		Topic:                    sub.Topic,
		AckDeadlineSeconds:       sub.AckDeadlineSeconds,
		MessageRetentionDuration: sub.MessageRetentionDuration,
	}
}

func (s Subscription) Title() string {
	parts := strings.Split(s.Name, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return s.Name
}

func (s Subscription) Subtitle() string {
	stateSymbol := "❌"
	if s.State == "ACTIVE" {
		stateSymbol = "✅"
	} else if s.State == "PENDING" {
		stateSymbol = "🕒"
	}

	topicParts := strings.Split(s.Topic, "/")
	topicName := s.Topic
	if len(topicParts) > 0 {
		topicName = topicParts[len(topicParts)-1]
	}

	return stateSymbol + " Topic: " + topicName +
		" | Ack Deadline: " + fmt.Sprintf("%ds", s.AckDeadlineSeconds) +
		" | Retention: " + s.MessageRetentionDuration
}

func (s Subscription) URL(config *gcloud.Config) string {
	return "https://console.cloud.google.com/cloudpubsub/subscription/detail/" + s.Title() + "?project=" + config.Project
}
