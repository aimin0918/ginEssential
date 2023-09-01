package models

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Users struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

func SelectUser(ctx context.Context, page, pageSize int) (users []Users, count int, err error) {
	model := Db.Model(&Users{})
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
	err = model.Order("id asc").Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}
