package user_service

import (
	"context"
	"ginessential/http/user_define"
	"ginessential/models"
)

func UserList(ctx context.Context, req user_define.UserListReq) (resp user_define.UserListResp, err error) {
	if req.Page <= 1 {
		req.Page = 1
	}

	if req.PageSize == 0 {
		req.PageSize = 20
	}
	list, count, err := models.UserList(ctx, req.Page, req.PageSize)
	if err != nil {
		return
	}
	resp.List = list
	resp.Count = count
	resp.Page = req.Page
	resp.PageSize = req.PageSize

	return
}

func GetUserDetail(ctx context.Context, Id int64) (resp user_define.UserDetailResp, err error) {
	user, err := models.GetUserDetail(ctx, Id)
	if err != nil {
		return
	}
	resp.User = user
	return
}

func UpsertUser(ctx context.Context, req user_define.UpsertUserResp) (err error) {
	User := models.Users{
		Id:        req.Id,
		Name:      req.Name,
		Telephone: req.Telephone,
		Password:  req.Password,
	}
	err = models.UpsertUser(ctx, User)
	if err != nil {
		return
	}
	return
}

func DeleteUser(ctx context.Context, id int64) (err error) {
	err = models.DeleteUser(ctx, id)
	if err != nil {
		return
	}
	return
}
