package main

import (
	simple "TikhampornSky/rabbitmq/watermill-direct"
	"context"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	wtmAmqp "github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	// Create an AMQP connection for publishing [OLD]
	commandPublisherOld := initPublisher("old")

	// Create an AMQP connection for publishing [NEW]
	commandPublisherNew := initPublisher("new")

	defer commandPublisherOld.Close()
	defer commandPublisherNew.Close()

	var myMockTime = time.Now().Add(time.Duration(0) * time.Hour)
	var mockTopicNew = "my-exchange"
	var mockTopicOld = "my-exchange-old"

	// Publish a message to new exchange to get delay (If needed) [NEW]
	msg := message.NewMessage(watermill.NewUUID(), []byte("***** This mock message is from: "+myMockTime.String()+" *****"))
	if simple.IsTimeClose(myMockTime) {
		log.Println("Time is close!!!")
		msg.Metadata.Set("x-delay", "10000")
	} else {
		log.Println("Time is not close")
	}

	if mockTopicNew == simple.AcceptedTopic {
		if err := commandPublisherNew.Publish(mockTopicNew, msg); err != nil {
			log.Fatalf("Failed to publish message: %s", err)
		} else {
			log.Printf("--> Published message: %s", msg.UUID)
		}
	} else {
		log.Println("Topic is not matched, do something else")
	}

	// Create an AMQP connection for subscribing [NEW]
	commandSubscriberNew, err := wtmAmqp.NewSubscriber(simple.AmqpConfigNew, simple.Logger)
	if err != nil {
		log.Panic("Cannot initialize subscriber", err)
	}
	defer commandSubscriberNew.Close()
	messages, err := commandSubscriberNew.Subscribe(context.Background(), "my-queue-tmp")
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}
	// =====================

	ch := make(chan int)
	go func() {
		for msg := range messages {
			log.Printf("--> [NEW] Received message: %s with UUID: %s", string(msg.Payload), msg.UUID)
			msg.Ack()
			log.Printf("[NEW] Acked message")

			if err := commandPublisherOld.Publish(mockTopicOld, msg); err != nil { // OLD
				log.Fatalf("Failed to publish message: %s", err)
			} else {
				log.Printf("--> [OLD] Published message: %s", msg.UUID)
				ch <- 1
			}
		}
	}()
	<-ch
}

func initPublisher(typePub string) *wtmAmqp.Publisher {
	var config wtmAmqp.Config
	if typePub == "old" {
		config = simple.AmqpConfigOld
	} else {
		config = simple.AmqpConfigNew
	}
	commandPublisher, err := wtmAmqp.NewPublisher(config, simple.Logger)
	if err != nil {
		log.Panic("Cannot initialize command publisher", err, " Type: ", typePub)
	}

	return commandPublisher
}
