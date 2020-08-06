package main

import (
	"github.com/gin-gonic/gin"
	"mwx563796/ginessential/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login",controller.Loginer)
	return r
}
