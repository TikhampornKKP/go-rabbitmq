package main

import (
	simple "TikhampornSky/rabbitmq/watermill-direct"
	"context"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	amqpSubscriber, err := amqp.NewSubscriber(simple.AmqpConfigNew, watermill.NewStdLogger(false, false))
	if err != nil {
		log.Fatalf("Failed to create AMQP subscriber: %s", err)
	}
	defer amqpSubscriber.Close()
	consumer, err := amqpSubscriber.Subscribe(context.Background(), "my-exchange")
	if err != nil {
		log.Fatalf("Failed to subscribe to exchange: %s", err)
	}

	amqpPublisher, err := amqp.NewPublisher(simple.AmqpConfigOld, watermill.NewStdLogger(false, false))
	if err != nil {
		log.Fatalf("Failed to create AMQP publisher: %s", err)
	}
	defer amqpPublisher.Close()

	var forever chan struct{}
	go func() {
		for msg := range consumer {
			log.Printf("--> [TMP] Received message: %s, %s", string(msg.Payload), msg.UUID)
			err := amqpPublisher.Publish("my-exchange-old", message.NewMessage(watermill.NewUUID(), msg.Payload))
			if err != nil {
				log.Printf("Failed to publish message to target exchange: %s", err)
			} else {
				log.Printf("--> [TMP] Published message to target exchange: %s", msg.UUID)
			}
		}
	}()
	<-forever
}
