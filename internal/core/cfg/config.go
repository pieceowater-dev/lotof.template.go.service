package cfg

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
)

var GossiperConf = gossiper.Config{
	Env: gossiper.EnvConfig{
		Required: []string{
			// indexes order is important for some getter functions
			"DATABASE_DSN", // DB connection string at 0 index
			"RABBITMQ_DSN", // RabbitMQ connection string at 1 index
		},
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
