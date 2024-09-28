package items

import (
	"application/internal/core/cfg"
	"application/internal/core/utils/common"
	"application/internal/pkg/items/ctrl"
	"application/internal/pkg/items/svc"
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
	service := svc.NewItemService(cfg.GetDB())
	controller := ctrl.NewItemController(service)

	var r any

	switch msg.Pattern {
	case CreateItem:
		r = controller.CreateItem(msg.Data)
	case FindAllItem:
		r = controller.GetItems(msg.Data)
	case FindOneItem:
		r = controller.GetItem(msg.Data)
	case UpdateItem:
		r = controller.UpdateItem(msg.Data)
	default:
		log.Println("Unknown action:", msg.Pattern)
		r = nil
		return "Unknown Items action"
	}

	defer gossiper.LogAction(msg.Pattern, common.ActionLog{Request: msg.Data, Response: r})
	return r
}
