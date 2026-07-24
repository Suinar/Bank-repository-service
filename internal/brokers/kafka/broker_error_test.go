package kafka

import (
	"context"
	"errors"
	"testing"

	kafkaGo "github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
)

func TestAdminEnsureTopicsRejectsInvalidInput(t *testing.T) {
	admin := &Admin{client: &topicCreatorStub{}}

	require.EqualError(t, admin.EnsureTopics(context.Background()), "at least one Kafka topic is required")
	require.EqualError(t, admin.EnsureTopics(context.Background(), ""), "Kafka topic name cannot be empty")
}

func TestAdminEnsureTopicsReturnsClientAndTopicErrors(t *testing.T) {
	clientErr := errors.New("broker unavailable")
	admin := &Admin{client: &topicCreatorStub{err: clientErr}}
	require.ErrorIs(t, admin.EnsureTopics(context.Background(), "bank.user.events"), clientErr)

	topicErr := errors.New("invalid topic")
	admin = &Admin{client: &topicCreatorStub{response: &kafkaGo.CreateTopicsResponse{
		Errors: map[string]error{"bank.user.events": topicErr},
	}}}
	require.ErrorIs(t, admin.EnsureTopics(context.Background(), "bank.user.events"), topicErr)
}
