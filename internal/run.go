package internal

import (
	"application/internal/core/cfg"
	"application/internal/pkg"
	"encoding/json"
	g "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

func Run() error {
	app := g.Bootstrap{}
	return g.Safely(func() {
		app.Setup(
			cfg.GossiperConf,
			func() any {
				defer g.DontPanic()
				log.Println("Binding db...")
				_ = g.Safely(func() {
					cfg.SetDB(app.PGDB.GetDB())
				})
				return nil
			},
			func(msg []byte) any {
				defer g.DontPanic()

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
	})

}
