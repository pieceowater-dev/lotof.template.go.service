package health

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

type Router struct {
	Patterns []string
}

func New() *Router {
	return &Router{
		Patterns: []string{"ping"},
	}
}

func (h *Router) HandleMessage(msg gossiper.AMQMessage) any {
	switch msg.Pattern {
	case "ping":
		log.Println("Received PING request")
		return "PONG"
	default:
		log.Println("Unknown action:", msg.Pattern)
		return "Unknown Health action"
	}
}
