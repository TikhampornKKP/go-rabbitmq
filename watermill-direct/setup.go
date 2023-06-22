package simple

import (
	"github.com/ThreeDotsLabs/watermill"
	wamqp "github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/rabbitmq/amqp091-go"
)

var (
	Logger        = watermill.NewStdLogger(false, false)
	AmqpConfigOld = wamqp.Config{
		Connection: wamqp.ConnectionConfig{
			AmqpURI: "amqp://guest:guest@localhost:5672/",
		},
		Marshaler: wamqp.DefaultMarshaler{},
		Exchange: wamqp.ExchangeConfig{
			GenerateName: func(topic string) string {
				return "my-exchange-old"
			},
			Type:    "fanout",
			Durable: true,
		},
		Queue: wamqp.QueueConfig{
			GenerateName: func(topic string) string {
				return topic
			},
			Durable: true,
		},
		QueueBind: wamqp.QueueBindConfig{
			GenerateRoutingKey: func(topic string) string {
				return ""
			},
		},
		Publish: wamqp.PublishConfig{
			GenerateRoutingKey: func(topic string) string {
				return ""
			},
		},
		Consume: wamqp.ConsumeConfig{
			Qos: wamqp.QosConfig{
				PrefetchCount: 1,
			},
		},
		TopologyBuilder: &wamqp.DefaultTopologyBuilder{},
	}

	AmqpConfigNew = wamqp.Config{
		Connection: wamqp.ConnectionConfig{
			AmqpURI: "amqp://guest:guest@localhost:5672/",
		},
		Marshaler: wamqp.DefaultMarshaler{},
		Exchange: wamqp.ExchangeConfig{
			GenerateName: func(topic string) string {
				return "my-exchange"
			},
			Type:    "x-delayed-message",
			Durable: true,
			Arguments: amqp091.Table{
				"x-delayed-type": "direct",
			},
		},
		Queue: wamqp.QueueConfig{
			GenerateName: func(topic string) string {
				return "my-queue-tmp"
			},
			Durable: true,
		},
		QueueBind: wamqp.QueueBindConfig{
			GenerateRoutingKey: func(topic string) string {
				return ""
			},
		},
		Publish: wamqp.PublishConfig{
			GenerateRoutingKey: func(topic string) string {
				return ""
			},
		},
		Consume: wamqp.ConsumeConfig{
			Qos: wamqp.QosConfig{
				PrefetchCount: 1,
			},
		},
		TopologyBuilder: &wamqp.DefaultTopologyBuilder{},
	}
)
