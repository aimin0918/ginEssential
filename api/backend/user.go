package backend

import (
	"ginessential/http/user_define"
	"ginessential/library/app"
	"ginessential/library/e"
	"github.com/gin-gonic/gin"

	"ginessential/service/user_service"
)

func UserList(c *gin.Context) {
	appG := app.Gin{C: c}
	req := user_define.UserListReq{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	user, err := user_service.UserList(c, req)
	if err != nil {
		appG.Response(e.ERROR, nil)
		return
	}
	appG.Response(e.SUCCESS, user)
	return

}

func GetUserDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	req := user_define.UserDetailRep{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	user, err := user_service.GetUserDetail(c, req.Id)
	if err != nil {
		appG.Response(e.ERROR, nil)
		return
	}
	appG.Response(e.SUCCESS, user)
	return
}

func UpsertUser(c *gin.Context) {
	appG := app.Gin{C: c}
	req := user_define.UpsertUserResp{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	err = user_service.UpsertUser(c, req)
	if err != nil {
		appG.Response(e.ERROR, nil)
		return

	}
	appG.Response(e.SUCCESS, err)
	return
}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	req := user_define.UserDetailRep{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	err = user_service.DeleteUser(c, req.Id)
	if err != nil {
		appG.Response(e.ERROR, nil)
		return
	}
	appG.Response(e.SUCCESS, err)
	return

}
