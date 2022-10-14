package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ipbproject/IPB-Vote/src/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.POST("/insert", controllers.Insert)
	r.GET("/get", controllers.Get)
	r.PUT("/update", controllers.Update)
	r.DELETE("/delete", controllers.Delete)
	r.Run()
}
