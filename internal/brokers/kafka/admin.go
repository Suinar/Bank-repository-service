package kafka

import (
	"context"
	"errors"
	"fmt"

	log "github.com/kVinsom/Bank-repository-service/internal/logging/broker/kafka"
	kafkaGo "github.com/segmentio/kafka-go"
)

const (
	topicPartitions        = 1
	topicReplicationFactor = 3
)

// Admin creates the Kafka topics owned by the repository service.
type Admin struct {
	client TopicCreator
}

// NewAdmin creates an admin client connected to the configured brokers.
func NewAdmin(brokers []string) *Admin {
	return &Admin{
		client: &kafkaGo.Client{Addr: kafkaGo.TCP(brokers...)},
	}
}

// EnsureTopics creates missing topics and accepts topics that already exist.
func (a *Admin) EnsureTopics(ctx context.Context, topics ...string) error {
	const operation = "ensure_topics"
	defer log.OperationStarted(operation)()

	if len(topics) == 0 {
		err := errors.New("at least one Kafka topic is required")
		log.DependencyFailed(operation, err)
		return err
	}

	configs := make([]kafkaGo.TopicConfig, 0, len(topics))
	for _, topic := range topics {
		if topic == "" {
			err := errors.New("Kafka topic name cannot be empty")
			log.DependencyFailed(operation, err)
			return err
		}
		configs = append(configs, kafkaGo.TopicConfig{
			Topic:             topic,
			NumPartitions:     topicPartitions,
			ReplicationFactor: topicReplicationFactor,
		})
	}

	response, err := a.client.CreateTopics(ctx, &kafkaGo.CreateTopicsRequest{Topics: configs})
	if err != nil {
		log.DependencyFailed(operation, err)
		return fmt.Errorf("create Kafka topics: %w", err)
	}
	for topic, topicErr := range response.Errors {
		if topicErr != nil && !errors.Is(topicErr, kafkaGo.TopicAlreadyExists) {
			log.DependencyFailed(operation, topicErr)
			return fmt.Errorf("create Kafka topic %q: %w", topic, topicErr)
		}
	}
	return nil
}
