package root_service

import (
	"context"
	"ginessential/http/root_define"
	"ginessential/models"
)

func RootList(ctx context.Context, req root_define.RootListReq) (resp root_define.RootListResp, err error) {
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	if req.Page <= 1 {
		req.Page = 1
	}

	list, count, err := models.RootList(ctx, req.Page, req.PageSize)
	if err != nil {
		return
	}
	resp.RootList = list
	resp.Count = count
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	return
}

func GetRootDetail(ctx context.Context, id int64) (resp root_define.RootDetailResp, err error) {
	root, err := models.GetRootById(ctx, id)
	if err != nil {
		return
	}
	resp.Root = root
	return
}

func UpsertRoot(ctx context.Context, req root_define.UpsertRootResp) (err error) {
	Root := models.Root{
		Id:       req.Id,
		RootName: req.RootName,
	}
	err = models.UpsertRoot(ctx, Root)
	if err != nil {
		return
	}
	return
}
