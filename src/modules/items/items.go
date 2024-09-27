package items

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"template/src/core/config"
	"template/src/modules/items/controllers"
	"template/src/modules/items/services"
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
	svc := services.NewItemService(config.GetDB())
	ctr := controllers.NewItemController(svc)
	switch msg.Pattern {
	case CreateItem:
		return ctr.CreateItem(msg.Data)
	case FindAllItem:
		return ctr.GetItems(msg.Data)
	case FindOneItem:
		return ctr.GetItem(msg.Data)
	case UpdateItem:
		return ctr.UpdateItem(msg.Data)
	default:
		log.Println("Unknown action:", msg.Pattern)
		return "Unknown Items action"
	}
}
