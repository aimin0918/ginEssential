package backend

import (
	"ginessential/http/root_define"
	"ginessential/library/app"
	"ginessential/library/e"
	"ginessential/service/root_service"
	"github.com/gin-gonic/gin"
)

func RootList(c *gin.Context) {
	appG := app.Gin{C: c}
	req := root_define.RootListReq{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	root, err := root_service.RootList(c, req)
	if err != nil {
		appG.Response(e.ERROR, nil)
		return
	}
	appG.Response(e.SUCCESS, root)
	return
}

func GetRootDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	req := root_define.RootDetailRep{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	root, err := root_service.GetRootDetail(c, req.Id)
	if err != nil {
		appG.Response(e.ERROR, nil)
		return
	}
	appG.Response(e.SUCCESS, root)
	return
}

func UpsertRoot(c *gin.Context) {
	appG := app.Gin{C: c}
	req := root_define.UpsertRootResp{}
	err := appG.GetRequest(&req)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, err.Error())
		return
	}
	err = root_service.UpsertRoot(c, req)
	if err != nil {
		appG.Response(e.ERROR, nil)
		return
	}
	appG.Response(e.SUCCESS, err)
	return
}
