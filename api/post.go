package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
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
func PostDetail(c *gin.Context) {
	postIdString := c.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string transform to int err:", err)
		tool.RespErrorWithData(c, "post_id格式有错")
		return
	}
	//根据postId得到post
	post, err := service.GetPostById(postId)
	if err != nil {
		fmt.Println("get post err:", err)
		tool.RespInternalError(c)
		return
	}
	//找到ID对应的评论
	comments, err := service.GetPostsComment(postId)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("get comments err:", err)
			tool.RespInternalError(c)
			return
		}
		return
	}
	var postDetail model.PostDetail
	postDetail.Post = post
	postDetail.Comments = comments
	tool.RespSuccessfulWithData(c, postDetail)
}
