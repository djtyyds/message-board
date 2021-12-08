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
		postGroup.POST("/", AddPost) //发布新留言
		postGroup.POST("/:post_id")  //修改留言

	}
}
