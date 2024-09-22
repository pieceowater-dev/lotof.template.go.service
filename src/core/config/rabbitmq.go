package config

import (
	"github.com/streadway/amqp"
)

var rabbitMQ *amqp.Connection

func GetRabbitMQConnection() *amqp.Connection {
	return rabbitMQ
}
