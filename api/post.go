package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"time"
)

func AddPost(c *gin.Context) {
	IUsername, _ := c.Get("username")
	username := IUsername.(string)
	txt := c.PostForm("txt")
	post := model.Post{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.AddPost(post)
	if err != nil {
		fmt.Println("add post err:", err)
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessful(c)
}
func BriefPost(c *gin.Context) {
	post, err := service.GetPosts()
	if err != nil {
		fmt.Println("get posts err:", err)
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfulWithData(c, post)
}
