package main

import (
	"application/internal/core/cfg"
	"application/internal/pkg"
	"encoding/json"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

//TODO: refactor handle filters/search/pagination etc
//TODO: response formats + error response format

//TODO: exception/panic handling
//TODO: less ANY
//TODO: add Docs & Comments

func main() {
	gossiper.Setup(
		cfg.GossiperConf,
		func() any {
			cfg.InitDB()
			return nil
		},
		func(msg []byte) any {
			var message gossiper.AMQMessage
			err := json.Unmarshal(msg, &message)
			if err != nil {
				log.Println("Failed to unmarshal custom message:", err)
				return nil
			}
			router := pkg.InitRouter()
			return router.HandleMessageRouter(message)
		},
	)
}
