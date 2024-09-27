package main

import (
	"encoding/json"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"template/src/core/config"
	"template/src/modules"
)

func main() {
	gossiper.Setup(config.GossiperConf, func(msg []byte) any {
		var message gossiper.AMQMessage
		err := json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Failed to unmarshal custom message:", err)
			return nil
		}
		return modules.HandleMessageRouter(message)
	})

	//db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	//if err != nil {
	//	log.Error(fmt.Errorf("failed to connect to database: %w", err), nil)
	//}
}
