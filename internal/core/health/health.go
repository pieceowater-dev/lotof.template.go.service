package health

import (
	g "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

const (
	Ping string = "ping"
)

type Router struct {
	Patterns []string
}

func New() *Router {
	return &Router{
		Patterns: []string{Ping},
	}
}

func (h *Router) HandleMessage(msg g.AMQMessage) any {
	switch msg.Pattern {
	case Ping:
		log.Println("Received PING request")
		return "PONG"
	default:
		log.Println("Unknown action:", msg.Pattern)
		return "Unknown Health action"
	}
}
