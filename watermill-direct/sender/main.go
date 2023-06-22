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

	const myMockTime = "11:30"
	var mockTopic = "my-exchange"

	// Publish a message
	msg := message.NewMessage(watermill.NewUUID(), []byte("This mock message is from "+myMockTime))
	if simple.IsTimeClose(myMockTime) {
		log.Println("Time is close")
		msg.Metadata.Set("x-delay", "5000")
	} else {
		log.Println("Time is not close")
	}

	if mockTopic == simple.AcceptedTopic {
		if err := commandPublisher.Publish(mockTopic, msg); err != nil {
			log.Fatalf("Failed to publish message: %s", err)
		} else {
			log.Printf("Published message: %s", msg.UUID)
		}
	} else {
		log.Println("Topic is not matched, do something else")
	}
}
