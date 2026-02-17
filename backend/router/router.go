package router

import (
	"github.com/gin-gonic/gin"
	"gestmoto/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/motoqueiro", controller.GetMotoqueiros)
	r.POST("/motoqueiro", controller.CreateMotoqueiro)
	r.GET("/motoqueiro/:id", controller.GetMotoqueiro)

	return r
}