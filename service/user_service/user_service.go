package user_service

import (
	"context"
	"ginessential/cache"
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

	resp.UserList = list
	resp.Count = count
	resp.Page = req.Page
	resp.PageSize = req.PageSize

	return
}

func GetUserDetail(ctx context.Context, Id int64) (resp user_define.UserDetailResp, err error) {
	uc := cache.UserCache{}
	user, err := uc.GetUserOrderById(ctx, Id)
	//user, err := models.GetUserDetail(ctx, Id)
	if err != nil {
		return
	}

	rc := cache.RootCache{}
	root, err := rc.GetRootOrder(ctx, user.RootId)
	//root, err := models.GetRootDetail(ctx, user.RootId)
	if err != nil {
		return
	}

	resp.User = user
	resp.Root = root
	return
}

func UpsertUser(ctx context.Context, req user_define.UpsertUserResp) (err error) {
	User := models.User{
		Id:        req.Id,
		Name:      req.Name,
		Telephone: req.Telephone,
		Password:  req.Password,
		RootId:    req.RootId,
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
