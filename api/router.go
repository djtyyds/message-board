package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	en := gin.Default()
	en.POST("/register", register) //注册
	en.POST("login", login)        //登录
	userGroup := en.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/ChangePassword", ChangePassword) //修改密码
	}
	postGroup := en.Group("/post")
	{
		postGroup.Use(auth)
		postGroup.POST("/", AddPost)           //发布新留言
		postGroup.POST("/:post_id")            //修改留言
		postGroup.GET("/", BriefPost)          //查看全部留言
		postGroup.GET("/:post_id", PostDetail) //查看一条留言详细信息和其下属评论
		postGroup.DELETE("/:post_id")          //删除留言

	}
	commentGroup := en.Group("/comment")
	{
		commentGroup.Use(auth)
		commentGroup.POST("/AddComment", AddCommit) //添加评论
		commentGroup.DELETE("/:comment_id")         //删除评论
	}
	en.Run()
}
