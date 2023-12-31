package main

import (
	v8 "ginessential/library/redis/v8"
	"ginessential/models"
	"ginessential/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

//go run main.go routes.go 启动
func main() {
	InitConfig()
	db := models.InitDB()
	defer db.Close()
	//创建一个默认的路由引擎
	r := gin.Default()
	r = routers.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) //启动web服务 默认8080
}

func InitConfig() {
	v8.LoadRedis()
	workDir, _ := os.Getwd()
	viper.SetConfigName("dev/db")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
