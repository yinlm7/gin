package main

import (
	"goAdmin/bootstrap"
	"goAdmin/global"
	"goAdmin/router"
)

func main() {
	//初始化配置文件
	bootstrap.InitializeConfig()
	//初始化redis
	bootstrap.InitRedis()
	//初始化mongodb连接池
	bootstrap.InitMongo()
	//注册路由
	router.InitRouter(global.App.Config.App.Port)
}
