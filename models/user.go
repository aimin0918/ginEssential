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
	BaseModel
}

func UserList(ctx context.Context, page, pageSize int) (users []Users, count int, err error) {
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

func GetUserDetail(ctx context.Context, Id int64) (users []Users, err error) {
	err = Db.Model(&Users{}).Where("id=?", Id).First(&users).Error
	return
}

func UpsertUser(ctx context.Context, users Users) (err error) {
	if users.Id == 0 {
		err = Db.Model(&Users{}).Create(&users).Error
	} else {
		err = Db.Model(&Users{}).Where("id = ?", users.Id).Updates(map[string]interface{}{
			"id":        users.Id,
			"name":      users.Name,
			"telephone": users.Telephone,
			"password":  users.Password,
		}).Error
	}
	if err != nil {
		return
	}
	return
}

func DeleteUser(ctx context.Context, id int64) (err error) {
	err = Db.Where("Id = ?", id).Delete(&Users{}).Error
	if err != nil {
		return
	}
	return
}
