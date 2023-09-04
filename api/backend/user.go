package backend

import (
	"ginessential/http/user_define"
	"ginessential/library/app"
	"ginessential/library/e"
	"github.com/gin-gonic/gin"

	"ginessential/service/user_service"
)

func SelectUser(c *gin.Context) {
	appG := app.Gin{C: c}
	req := user_define.UserListReq{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	user, err := user_service.SelectUser(c, req)
	if err != nil {
		appG.Response(e.ERROR, nil)
		return
	}
	appG.Response(e.SUCCESS, user)
	return

}
