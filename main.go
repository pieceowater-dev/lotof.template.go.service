package main

import (
	"application/internal/core/cfg"
	"application/internal/pkg"
	"encoding/json"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

//TODO: handle DTOs
//TODO: handle filters/search/pagination etc
//TODO: less ANY
//TODO: add Docs & Comments
//TODO: Dockerfile
//TODO: CI
//TODO: Readme.md

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
