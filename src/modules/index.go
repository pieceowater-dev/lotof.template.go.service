package modules

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"template/src/core/health"
	"template/src/utils/common"
)

type Router struct {
	modulePatterns map[string][]string
	modules        map[string]common.MessageHandler
}

func InitRouter() *Router {
	healthRouter := health.New()

	return &Router{
		modules: map[string]common.MessageHandler{
			"health": healthRouter.HandleMessage,
		},
		modulePatterns: map[string][]string{
			"health": healthRouter.Patterns,
		},
	}
}

func (r *Router) HandleMessageRouter(msg gossiper.AMQMessage) any {
	category := r.extractCategory(msg.Pattern)

	handler, exists := r.modules[category]
	if !exists {
		log.Println("No handler for category:", category)
		return "Handler not found"
	}
	return handler(msg)
}

func (r *Router) extractCategory(pattern string) string {
	for category, patterns := range r.modulePatterns {
		for _, p := range patterns {
			if pattern == p {
				return category
			}
		}
	}
	return "unknown"
}
