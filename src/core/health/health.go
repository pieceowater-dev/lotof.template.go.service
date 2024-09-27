package health

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

func HandleHealthMessage(msg gossiper.AMQMessage) any {
	switch msg.Pattern {
	case "ping":
		log.Println("Received PING request")
		return "PONG"

	default:
		log.Println("Unknown action:", msg.Pattern)
		return "Unknown Health action"
	}
}
