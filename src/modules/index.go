package modules

import (
	"github.com/gin-gonic/gin"
	"template/src/modules/items"
)

func Init(router *gin.Engine) {
	items.Init(router)
}
