package main

import (
	"article-manager/conf"
	"article-manager/router"
)

func main() {
	// 初始化配置
	conf.Init()

	// 初始化路由
	router.Init(conf.Conf)
	router.Stop()
}
