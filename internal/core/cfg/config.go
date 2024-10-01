package cfg

import (
	g "github.com/pieceowater-dev/lotof.lib.gossiper"
)

var GossiperConf = g.Config{
	Env: g.EnvConfig{
		Required: []string{
			// indexes order is important for some getter functions
			"DATABASE_DSN", // DB connection string at 0 index
			"RABBITMQ_DSN", // RabbitMQ connection string at 1 index
		},
	},
	AMQPConsumer: g.AMQPConsumerConfig{
		DSNEnv: "RABBITMQ_DSN",
		Queues: []g.QueueConfig{
			{
				Name:    "template_queue",
				Durable: true,
			},
		},
		Consume: []g.AMQPConsumeConfig{
			{
				Queue:    "template_queue",
				Consumer: "example_consumer",
				AutoAck:  true,
			},
		},
	},
}
