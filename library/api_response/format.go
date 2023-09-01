package api_response

import (
	"fmt"
	"net/http"
	"oceanlearn.teach/ginessential/library/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiResponse struct {
	ErroCode   int         `json:"errorCode"`
	ErrorText  string      `json:"errorText"`
	ReturnCode int         `json:"returnCode"`
	Data       interface{} `json:"data"`
}

type Empty struct {
}

func SendResult(ctx *gin.Context, code int, message string, data interface{}) {
	if message == "" {
		message = MessageMap[code]
	}

	returncode := ReturnErrorCode
	if code == CodeSuccess {
		returncode = ReturnSucessCode
	}
	if code != CodeSuccess {
		// 业务返回异常记录日志
		if ctx.Request.Method == "POST" {
			log.WarnWithCtx(ctx, fmt.Sprintf("%s %s failed", ctx.Request.Method, ctx.Request.URL), zap.Int("code", code), zap.String("message", message), zap.Any("data", data))
		}
	}
	// 不返回nil，返回nil的话mobile端json解析会认为有问题
	if data == nil {
		data = Empty{}
	}

	ctx.JSON(http.StatusOK, ApiResponse{ErroCode: code, ErrorText: message, Data: data, ReturnCode: returncode})
	return
}
