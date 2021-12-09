package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddComment(comment model.Comment) error {
	return dao.InsertComment(comment)
}
func GetPostsComment(PostId int) ([]model.Comment, error) {
	return dao.SelectCommentByPostId(PostId)
}
func GetCommentById(commentID int) (model.Comment, error) {
	return dao.SelectCommentById(commentID)
}
func AddCommentPraise(comment model.Comment) error {
	return dao.UpdateCommentPraise(comment)
}
