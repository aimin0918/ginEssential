package root_define

import "ginessential/models"

type RootListReq struct {
	Page     int `json:"page" form:"page"`           // 页码
	PageSize int `json:"page_size" form:"page_size"` // 条数
}

type RootListResp struct {
	RootList []models.Root `json:"root_list" form:"root_list"`
	Count    int           `json:"count" form:"count"`
	Page     int           `json:"page" form:"page"`
	PageSize int           `json:"page_size" form:"page_size"`
}

type RootDetailRep struct {
	Id int64 `json:"id" form:"id"`
}

type RootDetailResp struct {
	Root models.Root `json:"root" form:"root"`
}

type UpsertRootResp struct {
	Id       int64  `json:"id" form:"id"`
	RootName string `json:"root_name" form:"root_name"`
}
