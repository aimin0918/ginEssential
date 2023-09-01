package routers

import (
	"github.com/gin-gonic/gin"
	"oceanlearn.teach/ginessential/api/backend"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.GET("/api/auth/select", backend.SelectUser)

	return r
}
