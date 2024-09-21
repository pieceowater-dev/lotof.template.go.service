package items

import (
	"github.com/gin-gonic/gin"
	"template/src/core/config"
	"template/src/modules/items/controllers"
	"template/src/modules/items/routes"
	"template/src/modules/items/services"
)

func Init(router *gin.Engine) {
	db := config.GetDB()
	rabbitMQ := config.GetRabbitMQConnection()

	itemService := services.NewItemService(db, rabbitMQ)
	itemController := controllers.NewItemController(itemService)

	// Apply routes
	routes.ApplyRoutes(router, itemController)
}
