package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ipbproject/IPB-Vote/src/controllers"
)

func HanleRequest() {
	r := gin.Default()
	r.POST("/insert", controllers.Insert)
	r.Run()
}
