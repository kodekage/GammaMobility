package queue

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupId string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		GroupID:     groupId,
		Topic:       topic,
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset,
	})

	return &Consumer{Reader: r}
}

func (c *Consumer) ConsumeMessage(ctx context.Context) (kafka.Message, error) {
	return c.Reader.ReadMessage(ctx)
}

func (c *Consumer) Close() error {
	return c.Reader.Close()
}
