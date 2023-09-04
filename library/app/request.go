package app

import (
	"context"
	"errors"
	"ginessential/library/log"

	"github.com/astaxie/beego/validation"
	"go.uber.org/zap"
)

// MarkErrors logs error logs
func MarkErrors(ctx context.Context, errors []*validation.Error) {
	for _, err := range errors {
		log.ErrorWithCtx(ctx, err.Key, zap.String("msg", err.Message))
	}

	return
}

func (g *Gin) GetRequest(params interface{}) (err error) {
	err = g.C.ShouldBind(params)
	if err != nil {
		log.ErrorWithCtx(g.C, "参数错误", zap.Any("params", params), zap.Error(err))
		return
	}

	valid := validation.Validation{}
	ok, _ := valid.Valid(params)
	if !ok {
		MarkErrors(g.C, valid.Errors)
		return errors.New("参数校验失败")
	}

	return
}
