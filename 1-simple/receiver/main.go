package main

import (
	simple "TikhampornSky/rabbitmq/1-simple"
	"context"
	"log"

	wtmAmqp "github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

func main() {
	// Create an AMQP connection
	commandSubscriber, err := wtmAmqp.NewSubscriber(simple.AmqpConfig, simple.Logger)
	if err != nil {
		log.Panic("Cannot initialize subscriber", err)
	}
	defer commandSubscriber.Close()

	// Start consuming messages
	messages, err := commandSubscriber.Subscribe(context.Background(), "my-queue")
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	// Handle received messages
	var forever chan struct{}
	go func() {
		for msg := range messages {
			log.Printf("Received message: %s with UUID: %s", string(msg.Payload), msg.UUID)
			msg.Ack()
			log.Printf("Acked message")
		}
	}()
	<-forever
}
