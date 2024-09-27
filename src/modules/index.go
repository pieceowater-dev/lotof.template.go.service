package modules

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"template/src/core/health"
)

//docker run -d -p 3000:3000 --network tools_tools-network --name lotta.gtw -e MODE=dev -e RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672 ghcr.io/pieceowater-dev/lotof.template.gateway:6e471dc
//---items---
//createItem
//findAllItem
//findOneItem
//updateItem

type MessageHandler func(gossiper.AMQMessage) any

// Define category to patterns mapping
var categoryPatterns = map[string][]string{
	"health": {"ping"},
	// Add other categories and their patterns here
}

var handlers = map[string]MessageHandler{
	"health": health.HandleHealthMessage,
	// Add other handlers here
}

func HandleMessageRouter(msg gossiper.AMQMessage) any {
	category := extractCategory(msg.Pattern)

	handler, exists := handlers[category]
	if !exists {
		log.Println("No handler for category:", category)
		return "Handler not found"
	}
	return handler(msg)
}

func extractCategory(pattern string) string {
	for category, patterns := range categoryPatterns {
		for _, p := range patterns {
			if pattern == p {
				return category
			}
		}
	}
	return "unknown" // Default category if not found
}
