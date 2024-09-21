package main

import (
	"github.com/gin-gonic/gin"
	"template/src/core/config"
	"template/src/modules/items"
	log "template/src/utils/logs"
)

func main() {
	log.InitLogger()
	_, port, _, _ := config.Setup()
	router := gin.Default()

	items.Init(router)

	if err := router.Run(":" + port); err != nil {
		log.Error(err, nil)
	}
}
