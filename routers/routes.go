package routers

import (
	"ginessential/api/backend"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.GET("/api/auth/select", backend.SelectUser)

	return r
}
