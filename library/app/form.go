package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	e2 "oceanlearn.teach/ginessential/library/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e2.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e2.ERROR
	}
	if !check {
		MarkErrors(c, valid.Errors)
		return http.StatusBadRequest, e2.INVALID_PARAMS
	}

	return http.StatusOK, e2.SUCCESS
}
