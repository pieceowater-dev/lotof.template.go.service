package main

import (
	"application/internal/core/cfg"
	"application/internal/pkg"
	"encoding/json"
	g "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

//TODO: pack some shi to gossiper

//TODO: exception/panic handling
//TODO: refactor handle filters/search/pagination etc
//TODO: less ANY
//TODO: add Docs & Comments

func main() {
	g.Setup(
		cfg.GossiperConf,
		func() any {
			cfg.InitDB()
			return nil
		},
		func(msg []byte) any {
			var message g.AMQMessage
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
