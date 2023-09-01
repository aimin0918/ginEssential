package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	e2 "oceanlearn.teach/ginessential/library/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(errCode int, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  e2.GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) ResponseMessage(errCode int, errMsg string, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  errMsg,
		Data: data,
	})
	return
}
