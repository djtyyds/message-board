package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddPost(post model.Post) error {
	err := dao.AddPost(post)
	return err
}
func GetPosts() ([]model.Post, error) {
	return dao.SelectPost()
}
