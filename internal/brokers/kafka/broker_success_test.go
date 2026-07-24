package kafka

import (
	"context"
	"testing"

	kafkaGo "github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
)

type topicCreatorStub struct {
	response *kafkaGo.CreateTopicsResponse
	err      error
	request  *kafkaGo.CreateTopicsRequest
}

func (s *topicCreatorStub) CreateTopics(_ context.Context, request *kafkaGo.CreateTopicsRequest) (*kafkaGo.CreateTopicsResponse, error) {
	s.request = request
	return s.response, s.err
}

func TestAdminEnsureTopics(t *testing.T) {
	client := &topicCreatorStub{response: &kafkaGo.CreateTopicsResponse{}}
	admin := &Admin{client: client}

	err := admin.EnsureTopics(context.Background(), "bank.user.events", "bank.account.events")

	require.NoError(t, err)
	require.Len(t, client.request.Topics, 2)
	require.Equal(t, "bank.user.events", client.request.Topics[0].Topic)
	require.Equal(t, 1, client.request.Topics[0].NumPartitions)
	require.Equal(t, 3, client.request.Topics[0].ReplicationFactor)
}

func TestAdminEnsureTopicsAcceptsExistingTopic(t *testing.T) {
	client := &topicCreatorStub{response: &kafkaGo.CreateTopicsResponse{
		Errors: map[string]error{"bank.user.events": kafkaGo.TopicAlreadyExists},
	}}
	admin := &Admin{client: client}

	require.NoError(t, admin.EnsureTopics(context.Background(), "bank.user.events"))
}
