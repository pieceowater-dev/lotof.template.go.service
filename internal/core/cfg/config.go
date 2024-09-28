package cfg

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
)

var GossiperConf = gossiper.Config{
	Env: gossiper.EnvConfig{
		Required: []string{"DATABASE_DSN", "RABBITMQ_DSN"},
	},
	AMQPConsumer: gossiper.AMQPConsumerConfig{
		DSNEnv: "RABBITMQ_DSN",
		Queues: []gossiper.QueueConfig{
			{
				Name:    "template_queue",
				Durable: true,
			},
		},
		Consume: []gossiper.AMQPConsumeConfig{
			{
				Queue:    "template_queue",
				Consumer: "example_consumer",
				AutoAck:  true,
			},
		},
	},
}
