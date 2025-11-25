package queue

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

type Message struct {
	Value []byte
}

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(brokers...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		Async:                  true,
		RequiredAcks:           kafka.RequireOne,
	}

	return &Producer{Writer: w}
}

func (p *Producer) PublishMessage(ctx context.Context, key string, msg interface{}) error {
	messageBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	kafkaMsg := kafka.Message{
		Key:   []byte(key),
		Value: messageBytes,
		Time:  time.Now(),
	}

	return p.Writer.WriteMessages(ctx, kafkaMsg)
}
