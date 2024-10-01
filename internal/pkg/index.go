package pkg

import (
	"application/internal/core/health"
	"application/internal/core/utils/common"
	"application/internal/pkg/items"
	g "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
)

type Router struct {
	modules map[string]common.Module
}

func InitRouter() *Router {
	healthRouter := health.New()
	itemsRouter := items.New()

	return &Router{
		modules: map[string]common.Module{
			"health": {
				Patterns: healthRouter.Patterns,
				Handler:  healthRouter.HandleMessage,
			},
			"items": {
				Patterns: itemsRouter.Patterns,
				Handler:  itemsRouter.HandleMessage,
			},
		},
	}
}

func (r *Router) HandleMessageRouter(msg g.AMQMessage) any {
	category := r.extractCategory(msg.Pattern)

	module, exists := r.modules[category]
	if !exists {
		log.Println("No handler for category:", category)
		return "Handler not found"
	}
	return module.Handler(msg)
}

func (r *Router) extractCategory(pattern string) string {
	for category, module := range r.modules {
		for _, p := range module.Patterns {
			if pattern == p {
				return category
			}
		}
	}
	return "unknown"
}
