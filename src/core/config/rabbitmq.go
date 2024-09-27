package config

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"github.com/streadway/amqp"
)

var rabbitMQ *amqp.Connection

func GetRabbitMQConnection() *amqp.Connection {
	return rabbitMQ
}

func GetRabbitMQDSN() string {
	envInstance := &gossiper.Env{}
	val, err := envInstance.Get(GossiperConf.Env.Required[1])
	if err != nil {
		return ""
	}
	return val
}
