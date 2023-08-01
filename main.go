package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"oceanlearn.teach/ginessential/api"
	"oceanlearn.teach/ginessential/common"
	"os"
)

//go run main.go routes.go 启动
func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	//创建一个默认的路由引擎
	r := gin.Default()
	r = api.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) //启动web服务 默认8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
