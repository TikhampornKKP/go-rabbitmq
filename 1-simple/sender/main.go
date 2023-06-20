package main

import (
	simple "TikhampornSky/rabbitmq/1-simple"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	wtmAmqp "github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	// Create an AMQP connection
	commandPublisher, err := wtmAmqp.NewPublisher(simple.AmqpConfig, simple.Logger)
	if err != nil {
		log.Panic("Cannot initialize command publisher", err)
	}
	defer commandPublisher.Close()

	// Publish a message
	msg := message.NewMessage(watermill.NewUUID(), []byte("Hello1234"))
	if err := commandPublisher.Publish("my_topic", msg); err != nil {
		log.Fatalf("Failed to publish message: %s", err)
	} else {
		log.Printf("Published message: %s", msg.UUID)
	}
}
