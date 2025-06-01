package kafka

import (
	"context"
	"errors"
	"time"

	"github.com/segmentio/kafka-go"
)

func WithBrokers(brokers []string) OptsFunc {
	return func(k *Kafka) (err error) {
		k.brokers = brokers
		return
	}
}

func WithDefaultTopic(topic string) OptsFunc {
	return func(k *Kafka) (err error) {
		k.topic = topic
		return
	}
}

type Kafka struct {
	brokers []string
	topic   string
	writer  *kafka.Writer
}

type OptsFunc func(k *Kafka) error

func New(opts ...OptsFunc) (k *Kafka, err error) {
	k = &Kafka{}

	for _, opt := range opts {
		opt(k)
	}
	if len(k.brokers) == 0 {
		return nil, errors.New("brokers is required")
	}
	dialer := &kafka.Dialer{
		Timeout: 10 * time.Second,
	}

	k.writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  k.brokers,
		Balancer: &kafka.Hash{},
		Dialer:   dialer,
	})
	k.writer = kafka.NewWriter(kafka.WriterConfig{})
	return
}

func (k *Kafka) Close(ctx context.Context) (err error) {
	return k.writer.Close()
}
