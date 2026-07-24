package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/kVinsom/Bank-repository-service/internal/configs"
	log "github.com/kVinsom/Bank-repository-service/internal/logging/broker/kafka"
)

const (
	ensureAttempts   = 10
	ensureRetryDelay = 2 * time.Second
	ensureTimeout    = 5 * time.Second
)

// EnsureKafkaTopics connects to Kafka and creates the configured service topics.
func EnsureKafkaTopics(ctx context.Context, cfg *configs.Config) error {
	if !cfg.Kafka.Enabled {
		log.Disabled()
		return nil
	}

	admin := NewAdmin(cfg.Kafka.Brokers)
	var err error
	log.TopicsPreparing(cfg.Kafka.Topics)
	for attempt := 1; attempt <= ensureAttempts; attempt++ {
		attemptCtx, cancel := context.WithTimeout(ctx, ensureTimeout)
		err = admin.EnsureTopics(attemptCtx, cfg.Kafka.Topics...)
		cancel()
		if err == nil {
			log.TopicsReady(cfg.Kafka.Topics, attempt)
			return nil
		}
		if attempt == ensureAttempts {
			break
		}

		log.Retry(attempt, ensureAttempts, ensureRetryDelay, err)
		timer := time.NewTimer(ensureRetryDelay)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return ctx.Err()
		case <-timer.C:
		}
	}

	log.TopicsFailed(ensureAttempts, err)
	return fmt.Errorf("prepare Kafka topics after %d attempts: %w", ensureAttempts, err)
}
