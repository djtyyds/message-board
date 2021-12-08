package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func AddCommit(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)

	txt := c.PostForm("txt")
	postIdString := c.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err: ", err)
		tool.RespErrorWithData(c, "文章id有误")
		return
	}
	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		Username:    username,
		CommentTime: time.Now(),
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(c)
		return
	}

	tool.RespSuccessful(c)
}
