package models

import (
	"context"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	RootId    int64  `json:"root_id"`
	BaseModel
}

func UserList(ctx context.Context, page, pageSize int) (user []User, count int, err error) {
	model := Db.Model(&User{})
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
	err = model.Order("id asc").Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

func GetUserDetail(ctx context.Context, Id int64) (user User, err error) {
	err = Db.Model(&User{}).Where("id=?", Id).Order("id desc").Find(&user).Error
	return
}

func UpsertUser(ctx context.Context, user User) (err error) {
	if user.Id == 0 {
		err = Db.Model(&User{}).Create(&user).Error
	} else {
		err = Db.Model(&User{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
			"id":        user.Id,
			"name":      user.Name,
			"telephone": user.Telephone,
			"password":  user.Password,
			"root_id":   user.RootId,
		}).Error
	}
	if err != nil {
		return
	}
	return
}

func DeleteUser(ctx context.Context, id int64) (err error) {
	err = Db.Where("Id = ?", id).Delete(&User{}).Error
	if err != nil {
		return
	}
	return
}
