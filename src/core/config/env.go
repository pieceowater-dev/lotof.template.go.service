package config

import (
	"template/src/utils/env"
	log "template/src/utils/logs"
)

func SetupEnv() {
	env.LoadEnv()

	requiredVars := []string{"RABBITMQ_URL", "DATABASE_URL"}
	if err := env.CheckRequiredEnv(requiredVars); err != nil {
		log.Error(err, nil)
	}
}
