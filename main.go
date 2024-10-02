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
//TODO: less ANY

func main() {
	app := g.Bootstrap{}
	app.Setup(
		cfg.GossiperConf,
		func() any {
			log.Println("Binding db...")
			cfg.SetDB(app.DB.GetDB())
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
