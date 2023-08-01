package api

import (
	"github.com/gin-gonic/gin"
	"oceanlearn.teach/ginessential/middleware"
	"oceanlearn.teach/ginessential/model/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/auth/delete", controller.Delete)
	r.POST("/api/auth/update", controller.Update)

	return r
}
