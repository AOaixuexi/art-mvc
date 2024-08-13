package router

import "article-manager/conf"

func Init(c *conf.Config) {
	StartHttp(c)
}

func Stop() {
	srv.Close()
}
