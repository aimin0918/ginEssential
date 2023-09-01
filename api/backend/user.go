package backend

import (
	"github.com/gin-gonic/gin"
	"oceanlearn.teach/ginessential/http/user_define"
	app2 "oceanlearn.teach/ginessential/library/app"
	e2 "oceanlearn.teach/ginessential/library/e"

	"oceanlearn.teach/ginessential/service/user_service"
)

func SelectUser(c *gin.Context) {
	appG := app2.Gin{C: c}
	req := user_define.UserListReq{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e2.INVALID_PARAMS, err.Error())
		return
	}
	user, err := user_service.SelectUser(c, req)
	if err != nil {
		appG.Response(e2.ERROR, nil)
		return
	}
	appG.Response(e2.SUCCESS, user)
	return

}
