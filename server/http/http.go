package http

import (
	"preh/config"

	"github.com/gin-gonic/gin"
)

var (
	// SetupUserApi = setupUserApi
	InitEngine = initEngine
)

var router *gin.Engine

func initEngine() {
	gin.SetMode(config.GetGinMode())
	router = gin.Default()

	api := router.Group("/api")
	//test
	api.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "test successful",
		})
	})

	//user
	api.GET("User/:id", GetUserById)
	api.POST("User", CreateUser)
	api.GET("Order/:id", GetOrderById)
	api.POST("Order", CreateOrder)
	api.PATCH("Order/:id", UpdateOrderById)

	router.Run(config.GetGinUrl())
}
