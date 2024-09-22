package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"template/src/utils/env"
	log "template/src/utils/logs"
)

func Setup() (string, string, *gorm.DB, *amqp.Connection) {
	SetupEnv()

	port := env.GetEnvOrDefault("PORT", "3003")
	mode := env.GetEnvOrDefault("MODE", "dev")

	if mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

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

	return port, mode, db, rabbitmqConn
}
