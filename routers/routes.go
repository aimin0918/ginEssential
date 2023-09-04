package routers

import (
	"ginessential/api/backend"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.GET("/api/auth/userList", backend.UserList)           //用户列表
	r.GET("/api/auth/getUserDetail", backend.GetUserDetail) //用户详情
	r.POST("/api/auth/upsertUser", backend.UpsertUser)      //添加修改用户
	r.GET("/api/auth/deleteUser", backend.DeleteUser)       //删除会员

	return r
}
