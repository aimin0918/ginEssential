package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"oceanlearn.teach/ginessential/common"
)



func main(){
	db := common.GetDB()
	defer db.Close()
	//创建一个默认的路由引擎
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())//启动web服务 默认8080
}



