package main

import (
	"go-demo/configs"
	"go-demo/pkg/model"
	"go-demo/pkg/routers"
)

func main() {
	// 加载配置文件
	configs.InitConf()
	// 加载数据库驱动
	model.ConnectDatabase()
	defer model.CloseDB()
	// 注册路由并启动web
	routers.InitRouter()
}
