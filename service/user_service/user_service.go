package user_service

import (
	"context"
	"oceanlearn.teach/ginessential/http/user_define"
	"oceanlearn.teach/ginessential/models"
)

func SelectUser(ctx context.Context, req user_define.UserListReq) (resp user_define.UserListResp, err error) {
	if req.Page <= 1 {
		req.Page = 1
	}

	if req.PageSize == 0 {
		req.PageSize = 20
	}
	list, count, err := models.SelectUser(ctx, req.Page, req.PageSize)
	if err != nil {
		return
	}
	resp.List = list
	resp.Count = count
	resp.Page = req.Page
	resp.PageSize = req.PageSize

	return
}
