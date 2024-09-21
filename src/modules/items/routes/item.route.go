package routes

import (
	"github.com/gin-gonic/gin"
	"template/src/modules/items/controllers"
)

func ApplyRoutes(r *gin.Engine, itemController *controllers.ItemController) {
	r.POST("/items", itemController.CreateItem)
	r.GET("/items/:id", itemController.GetItem)
	r.PUT("/items/:id", itemController.UpdateItem)
	r.DELETE("/items/:id", itemController.DeleteItem)
}
