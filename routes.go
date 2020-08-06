package main

import (
	"github.com/gin-gonic/gin"
	"mwx563796/ginessential/controller"
	"mwx563796/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login",controller.Loginer)
	r.GET("/api/auth/info",middleware.AuthMiddleware(),controller.Info)
	return r
}
