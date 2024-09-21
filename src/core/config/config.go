package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"template/src/utils/env"
	log "template/src/utils/logs"
)

var db *gorm.DB
var rabbitMQ *amqp.Connection

func Setup() (string, string, *gorm.DB, *amqp.Connection) {
	SetupEnv()

	port := env.GetEnvOrDefault("PORT", "3003")
	mode := env.GetEnvOrDefault("MODE", "dev")

	dbURL, _ := env.GetEnv("DATABASE_URL", true)
	rabbitmqURL, _ := env.GetEnv("RABBITMQ_URL", true)

	log.Info("Environment configured:", map[string]interface{}{
		"MODE": mode,
		"PORT": port,
		"DB":   dbURL,
		"MQ":   rabbitmqURL,
	})

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Error(fmt.Errorf("failed to connect to database: %w", err), nil)
	}

	rabbitmqConn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Error(fmt.Errorf("failed to connect to RabbitMQ: %w", err), nil)
	}

	CheckDBConnection(dbURL)
	CheckRabbitMQConnection(rabbitmqURL)

	return mode, port, db, rabbitmqConn
}

func GetDB() *gorm.DB {
	return db
}

func GetRabbitMQConnection() *amqp.Connection {
	return rabbitMQ
}
