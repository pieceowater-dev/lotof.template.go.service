package main

import (
	"encoding/json"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"template/src/core/config"
	"template/src/modules"
)

//TODO: handle DTOs
//TODO: handle filters/search/pagination etc
//TODO: less ANY
//TODO: add Docs & Comments

func main() {
	gossiper.Setup(config.GossiperConf,
		func() any {
			config.InitDB()
			return nil
		},
		func(msg []byte) any {
			var message gossiper.AMQMessage
			err := json.Unmarshal(msg, &message)
			if err != nil {
				log.Println("Failed to unmarshal custom message:", err)
				return nil
			}
			router := modules.InitRouter()
			return router.HandleMessageRouter(message)
		},
	)
}
