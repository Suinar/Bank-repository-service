package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/kVinsom/Bank-repository-service/internal/logging/broker/kafka"
	kafkaGo "github.com/segmentio/kafka-go"
)

const outboxSchema = `
CREATE TABLE IF NOT EXISTS outbox_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    topic TEXT NOT NULL,
    event_type TEXT NOT NULL,
    aggregate_type TEXT NOT NULL,
    aggregate_id TEXT NOT NULL,
    payload JSONB NOT NULL,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    published_at TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_outbox_events_pending
    ON outbox_events (occurred_at) WHERE published_at IS NULL;

CREATE OR REPLACE FUNCTION enqueue_domain_event() RETURNS TRIGGER AS $$
DECLARE
    row_data JSONB;
    entity TEXT;
    operation TEXT;
BEGIN
    row_data := CASE WHEN TG_OP = 'DELETE' THEN to_jsonb(OLD) ELSE to_jsonb(NEW) END;
    row_data := row_data - 'password_hash';
    entity := CASE TG_TABLE_NAME
        WHEN 'users' THEN 'user'
        WHEN 'currencies' THEN 'currency'
        WHEN 'accounts' THEN 'account'
        WHEN 'cards' THEN 'card'
        WHEN 'credits' THEN 'credit'
        WHEN 'deposits' THEN 'deposit'
    END;
    operation := CASE TG_OP
        WHEN 'INSERT' THEN 'created'
        WHEN 'UPDATE' THEN 'updated'
        WHEN 'DELETE' THEN 'deleted'
    END;
    INSERT INTO outbox_events(topic, event_type, aggregate_type, aggregate_id, payload)
    VALUES ('bank.' || entity || '.events', entity || '.' || operation, entity, row_data->>'id', row_data);
    RETURN CASE WHEN TG_OP = 'DELETE' THEN OLD ELSE NEW END;
END;
$$ LANGUAGE plpgsql;

DO $$
DECLARE table_name TEXT;
BEGIN
    FOREACH table_name IN ARRAY ARRAY['users','currencies','accounts','cards','credits','deposits']
    LOOP
        EXECUTE format('DROP TRIGGER IF EXISTS trg_%s_outbox ON %I', table_name, table_name);
        EXECUTE format(
            'CREATE TRIGGER trg_%s_outbox AFTER INSERT OR UPDATE OR DELETE ON %I FOR EACH ROW EXECUTE FUNCTION enqueue_domain_event()',
            table_name, table_name
        );
    END LOOP;
END $$;`

type outboxEvent struct {
	ID            string          `db:"id" json:"event_id"`
	Topic         string          `db:"topic" json:"-"`
	EventType     string          `db:"event_type" json:"event_type"`
	AggregateType string          `db:"aggregate_type" json:"aggregate_type"`
	AggregateID   string          `db:"aggregate_id" json:"aggregate_id"`
	Payload       json.RawMessage `db:"payload" json:"payload"`
	OccurredAt    time.Time       `db:"occurred_at" json:"occurred_at"`
	EventVersion  int             `db:"-" json:"event_version"`
	Producer      string          `db:"-" json:"producer"`
}

// EnsureOutboxSchema installs the transactional outbox and table triggers.
func EnsureOutboxSchema(ctx context.Context, db *sqlx.DB) error {
	if _, err := db.ExecContext(ctx, outboxSchema); err != nil {
		return fmt.Errorf("ensure outbox schema: %w", err)
	}
	return nil
}

// RunOutboxPublisher publishes committed outbox rows until ctx is cancelled.
func RunOutboxPublisher(ctx context.Context, db *sqlx.DB, brokers []string, clientID string) error {
	writer := &kafkaGo.Writer{
		Addr:         kafkaGo.TCP(brokers...),
		Balancer:     &kafkaGo.Hash{},
		RequiredAcks: kafkaGo.RequireAll,
		BatchTimeout: 100 * time.Millisecond,
	}
	defer writer.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		if err := publishBatch(ctx, db, writer, clientID); err != nil && ctx.Err() == nil {
			log.OutboxPublishFailed(err)
		}
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
	}
}

func publishBatch(ctx context.Context, db *sqlx.DB, writer *kafkaGo.Writer, producer string) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var events []outboxEvent
	if err := tx.SelectContext(ctx, &events, `
SELECT id::text, topic, event_type, aggregate_type, aggregate_id, payload, occurred_at
FROM outbox_events
WHERE published_at IS NULL
ORDER BY occurred_at
LIMIT 100
FOR UPDATE SKIP LOCKED`); err != nil {
		return err
	}
	for i := range events {
		events[i].EventVersion = 1
		events[i].Producer = producer
		value, err := json.Marshal(events[i])
		if err != nil {
			return err
		}
		if err := writer.WriteMessages(ctx, kafkaGo.Message{
			Topic: events[i].Topic,
			Key:   []byte(events[i].AggregateID),
			Value: value,
			Time:  events[i].OccurredAt,
		}); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `UPDATE outbox_events SET published_at = NOW() WHERE id = $1`, events[i].ID); err != nil {
			return err
		}
	}
	return tx.Commit()
}
