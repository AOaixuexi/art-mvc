package router

import (
	"article-manager/conf"
	"article-manager/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	srv *service.Service
)

func StartHttp(c *conf.Config) {
	srv = service.New(c)

	r := gin.Default()
	r.Use(cors.Default())

	// 测试
	r.GET("/health", health)

	// 用户登陆注册
	r.POST("/register", srv.CreateUser)
	r.POST("/login", srv.FindAUser)

	// 用户界面操作
	paperGroup := r.Group("/user/:user_id")
	{
		// 增添文章
		paperGroup.POST("/", srv.CreatePaper)

		// 展示该用户所有文章
		paperGroup.GET("/", srv.GetPaperList)

		// 展示该用户文章，通过文章id或文章名
		paperGroup.GET("/:param", srv.GetPaper)

		// 更新文章
		paperGroup.PUT("/:id", srv.UpdateAPaper)

		// 删除文章
		paperGroup.DELETE("/:id", srv.DeleteAPaper)

		// 查找文章
		paperGroup.GET("/paper_search", srv.GetPapers)

		// 评论文章
		paperGroup.POST("/paper/:paper_id/comment", srv.CreateComment)

		// 查看文章评论
		paperGroup.GET("/paper/:paper_id/comment", srv.GetComments)

		// 删除评论
		paperGroup.DELETE("/paper/:paper_id/comment/:comment_id", srv.DeleteComment)

		// 添加子评论
		paperGroup.POST("/comment/:comment_id/subcomment", srv.CreateSubComment)

		// 查看所有子评论
		paperGroup.GET("/comment/:comment_id/subcomment", srv.GetSubComments)

		// 删除子评论
		paperGroup.DELETE("/comment/:comment_id/subcomment/:subcomment_id", srv.DeleteSubComment)
	}

	r.Run(c.HttpServer.Addr)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
