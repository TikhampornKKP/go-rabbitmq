package main

import (
	"context"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	amqpConfig := amqp.NewDurableQueueConfig("my-queue")

	amqpSubscriber, err := amqp.NewSubscriber(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		log.Fatalf("Failed to create AMQP subscriber: %s", err)
	}

	amqpPublisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		log.Fatalf("Failed to create AMQP publisher: %s", err)
	}

	consumer, err := amqpSubscriber.Subscribe(context.Background(), "my-exchange")
	if err != nil {
		log.Fatalf("Failed to subscribe to exchange: %s", err)
	}

	go func() {
		for msg := range consumer {
			log.Printf("Received message: %s", string(msg.Payload))
			err := amqpPublisher.Publish("my-exchange-old", message.NewMessage(watermill.NewUUID(), msg.Payload))
			if err != nil {
				log.Printf("Failed to publish message to target exchange: %s", err)
			} else {
				log.Printf("Published message to target exchange")
			}
		}
	}()
}
