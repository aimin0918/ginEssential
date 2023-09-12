package models

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Root struct {
	Id       int64  `json:"id"`
	RootName string `json:"root_name"`
	BaseModel
}

func RootList(ctx context.Context, page, pageSize int) (root []Root, count int, err error) {
	model := Db.Model(&Root{})
	err = model.Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		return
	}
	if pageSize > 0 {
		model = model.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	err = model.Order("id asc").Find(&root).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

func GetRootById(ctx context.Context, Id int64) (root Root, err error) {
	err = Db.Model(&Root{}).Where("id=?", Id).Order("id desc").First(&root).Error
	return
}

func GetRootDetail(ctx context.Context, Id int64) (root Root, err error) {
	err = Db.Model(&Root{}).Where("id=?", Id).Order("id desc").Find(&root).Error
	return
}

func UpsertRoot(ctx context.Context, root Root) (err error) {
	if root.Id == 0 {
		err = Db.Model(&Root{}).Create(&root).Error
	} else {
		err = Db.Model(&User{}).Where("id = ?", root.Id).Updates(map[string]interface{}{
			"id":        root.Id,
			"root_name": root.RootName,
		}).Error
	}
	if err != nil {
		return
	}
	return
}
