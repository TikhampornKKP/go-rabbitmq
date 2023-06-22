package main

import (
	simple "TikhampornSky/rabbitmq/watermill-direct"
	"context"
	"log"

	wtmAmqp "github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

func main() {
	// Create an AMQP connection for publishing
	commandPublisher, err := wtmAmqp.NewPublisher(simple.AmqpConfigOld, simple.Logger)
	if err != nil {
		log.Panic("Cannot initialize command publisher", err)
	}
	defer commandPublisher.Close()

	var mockTopic = "my-exchange-old"

	// Create an AMQP connection for subscribing
	commandSubscriber, err := wtmAmqp.NewSubscriber(simple.AmqpConfigNew, simple.Logger)
	if err != nil {
		log.Panic("Cannot initialize subscriber", err)
	}
	defer commandSubscriber.Close()
	messages, err := commandSubscriber.Subscribe(context.Background(), "my-queue-tmp")
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	var forever chan struct{}
	go func() {
		for msg := range messages {
			log.Printf("--> [OLD] Received message: %s with UUID: %s", string(msg.Payload), msg.UUID)
			msg.Ack()
			log.Printf("[OLD] Acked message")

			if err := commandPublisher.Publish(mockTopic, msg); err != nil {
				log.Fatalf("Failed to publish message: %s", err)
			} else {
				log.Printf("--> [OLD] Published message: %s", msg.UUID)
			}
		}
	}()
	<-forever
}
