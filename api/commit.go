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
func LikeComment(c *gin.Context) {
	PostDetail(c)
	iCommentId, _ := c.Get("commentId")
	CommentId := iCommentId.(int)
	//根据CommentId得到comment
	comment, err := service.GetCommentById(CommentId)
	if err != nil {
		fmt.Println("get comment err:", err)
		tool.RespInternalError(c)
		return
	}
	err = service.AddCommentPraise(comment)
	if err != nil {
		fmt.Println("add praise err :", err)
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessful(c)
}
func AddCommentFromComment(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	iCommentId, _ := c.Get("comment_id")
	commentId := iCommentId.(int)
	txt := c.PostForm("txt")
	addComment := model.Comment{
		Id:          commentId,
		Txt:         txt,
		Username:    username,
		CommentTime: time.Now(),
	}
	comment, err := service.GetCommentById(commentId)
	root := model.Node{
		Left: nil,
		Data: comment,
	}
	err = service.AddCommentFromComment(commentId, comment)
	if err != nil {
		fmt.Println("add comment err:", err)
		tool.RespInternalError(c)
	}
	a := model.Node{
		Left: nil,
		Data: addComment,
	}
	root.Left = &a
	tool.RespSuccessful(c)
}
func ShowCommentTrees(c *gin.Context) {

}
