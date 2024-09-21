package config

import (
	"fmt"
	"github.com/streadway/amqp"
	log "template/src/utils/logs"
)

var rabbitMQ *amqp.Connection

func GetRabbitMQConnection() *amqp.Connection {
	return rabbitMQ
}

func CheckRabbitMQConnection(rabbitmqURL string) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Error(fmt.Errorf("failed to connect to RabbitMQ: %w", err), nil)
		return
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	log.Info("Connected to RabbitMQ", map[string]interface{}{
		"status": "connected",
		"url":    rabbitmqURL,
	})
}
