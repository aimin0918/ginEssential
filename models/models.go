package models

import (
	"fmt"
	"ginessential/utils"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var Db *gorm.DB

type BaseModel struct {
	CreatedAt *utils.XTime `json:"created_at"`
	UpdatedAt *utils.XTime `json:"updated_at"`
}

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	//db.AutoMigrate(&models.Users{})

	Db = db
	return db
}

func GetDB() *gorm.DB {
	return Db
}
