package main

import (
	"github.com/gin-gonic/gin"
	"template/src/core/config"
	"template/src/modules"
	log "template/src/utils/logs"
)

func main() {
	log.InitLogger()
	port, _, _, _ := config.Setup()
	router := gin.Default()

	modules.Init(router)

	if err := router.Run(":" + port); err != nil {
		log.Error(err, nil)
	}
}
