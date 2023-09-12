package user_define

import "ginessential/models"

type UserListReq struct {
	Page     int `json:"page" form:"page"`           // 页码
	PageSize int `json:"page_size" form:"page_size"` // 条数
}

type UserListResp struct {
	UserList []models.User `json:"user_list" form:"user_list"`
	Count    int           `json:"count" form:"count"`
	Page     int           `json:"page" form:"page"`
	PageSize int           `json:"page_size" form:"page_size"`
}

type UserDetailRep struct {
	Id int64 `json:"id" form:"id"`
}

type UserDetailResp struct {
	User models.User `json:"user" form:"user"`
	Root models.Root `json:"root" form:"root"`
}

type UpsertUserResp struct {
	Id        int64  `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Telephone string `json:"telephone" form:"telephone"`
	Password  string `json:"password" form:"password"`
	RootId    int64  `json:"root_id" form:"root_id"`
}
