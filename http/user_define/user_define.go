package user_define

import "ginessential/models"

type Users struct {
	Id        int64  `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Telephone string `json:"telephone" form:"telephone"`
	Password  string `json:"password" form:"password"`
}

type UserListReq struct {
	Page     int `json:"page" form:"page"`           // 页码
	PageSize int `json:"page_size" form:"page_size"` // 条数
}

type UserListResp struct {
	List     []models.Users `json:"list" form:"list"`
	Count    int            `json:"count" form:"count"`
	Page     int            `json:"page" form:"page"`
	PageSize int            `json:"page_size" form:"page_size"`
}
