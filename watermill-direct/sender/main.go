package main

import (
	simple "TikhampornSky/rabbitmq/watermill-direct"
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

	// Mock Data
	const myMockTime = "9:30"
	var myTopic = "my-exchange&" + myMockTime
	// ================================

	// Publish a message
	msg := message.NewMessage(watermill.NewUUID(), []byte("This mock message is from "+myMockTime))
	msg.Metadata.Set("x-delay", "5000")

	if err := commandPublisher.Publish(myTopic, msg); err != nil {
		log.Fatalf("Failed to publish message: %s", err)
	} else {
		log.Printf("Published message: %s", msg.UUID)
	}
}
