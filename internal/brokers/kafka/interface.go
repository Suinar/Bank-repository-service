package kafka

import (
	"context"

	kafkaGo "github.com/segmentio/kafka-go"
)

// TopicCreator is the Kafka operation required by Admin.
type TopicCreator interface {
	CreateTopics(context.Context, *kafkaGo.CreateTopicsRequest) (*kafkaGo.CreateTopicsResponse, error)
}
