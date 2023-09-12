package routers

import (
	"ginessential/api/backend"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.GET("/api/auth/userList", backend.UserList)           // 用户列表
	r.GET("/api/auth/getUserDetail", backend.GetUserDetail) // 用户详情
	r.POST("/api/auth/upsertUser", backend.UpsertUser)      // 编辑用户
	r.GET("/api/auth/deleteUser", backend.DeleteUser)       // 删除会员

	r.GET("/api/auth/rootList", backend.RootList)           // 权限列表
	r.GET("/api/auth/getRootDetail", backend.GetRootDetail) // 权限详情
	r.GET("/api/auth/upsertRoot", backend.UpsertRoot)       // 编辑权限

	return r
}
