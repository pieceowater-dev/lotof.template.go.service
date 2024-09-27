package items

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

const (
	CreateItem  string = "createItem"
	FindAllItem string = "findAllItem"
	FindOneItem string = "findOneItem"
	UpdateItem  string = "updateItem"
)

type Router struct {
	Patterns []string
}

func New() *Router {
	return &Router{
		Patterns: []string{
			CreateItem,
			FindAllItem,
			FindOneItem,
			UpdateItem,
		},
	}
}

func (h *Router) HandleMessage(msg gossiper.AMQMessage) any {
	switch msg.Pattern {
	case CreateItem:
		log.Println("Received createItem request")
		return "createItem"
	case FindAllItem:
		log.Println("Received findAllItem request")
		return "findAllItem"
	case FindOneItem:
		log.Println("Received findOneItem request")
		return "findOneItem"
	case UpdateItem:
		log.Println("Received updateItem request")
		return "updateItem"
	default:
		log.Println("Unknown action:", msg.Pattern)
		return "Unknown Items action"
	}
}
