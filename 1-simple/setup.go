package simple

import (
	"github.com/ThreeDotsLabs/watermill"
	wtmAmqp "github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

var (
	Logger     = watermill.NewStdLogger(false, false)
	AmqpConfig = wtmAmqp.NewDurableQueueConfig("amqp://guest:guest@localhost:5672/")
)
