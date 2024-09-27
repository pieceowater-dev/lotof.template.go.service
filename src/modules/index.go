package modules

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"template/src/core/health"
	"template/src/utils/common"
)

type Router struct {
	modules map[string]common.Module
}

func InitRouter() *Router {
	healthRouter := health.New()

	return &Router{
		modules: map[string]common.Module{
			"health": {
				Patterns: healthRouter.Patterns,
				Handler:  healthRouter.HandleMessage,
			},
		},
	}
}

func (r *Router) HandleMessageRouter(msg gossiper.AMQMessage) any {
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
