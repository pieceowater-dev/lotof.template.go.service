package config

import (
	"template/src/utils/env"
	log "template/src/utils/logs"
)

func Setup() {
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

	CheckDBConnection(dbURL)
	CheckRabbitMQConnection(rabbitmqURL)
}
