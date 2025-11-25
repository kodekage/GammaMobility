package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/kodekage/gamma_mobility/dto"
	"github.com/kodekage/gamma_mobility/internal/logger"
	"github.com/kodekage/gamma_mobility/internal/queue"
	"github.com/kodekage/gamma_mobility/utils"
	"github.com/segmentio/kafka-go"
)

var (
	redisClient = utils.RedisClient()
)

func main() {
	consumer := queue.NewConsumer([]string{"localhost:9092"}, "payments", "payment-workers")

	logger.Info("Payment Worker started, waiting for messages...")

	for {
		msg, err := consumer.ConsumeMessage(context.Background())
		if err != nil {
			logger.Error("Error consuming message: " + err.Error())
			time.Sleep(500 * time.Millisecond)
			continue
		}

		go processMessage(msg)
	}
}

func processMessage(msg interface{}) {
	m := msg.(kafka.Message)

	var payload dto.CreateCustomerPaymentRequest
	if err := json.Unmarshal(m.Value, &payload); err != nil {
		logger.Error("Error unmarshaling message: " + err.Error())
		return
	}

	ctx := context.Background()
	idempotencyKey := "payment:" + payload.TransactionReference

	ok, _ := redisClient.SetNX(ctx, idempotencyKey, "1", 24*time.Hour).Result()
	if !ok {
		logger.Info("Duplicate payment detected, skipping: " + payload.TransactionReference)
		return
	}

	logger.Info("Processing payment for TransactionReference: " + payload.TransactionReference)
}
