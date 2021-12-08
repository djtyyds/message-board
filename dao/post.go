package dao

import "message-board/model"

func AddPost(post model.Post) error {
	_, err := DB.Exec("INSERT INTO post(username, txt, post_time, update_time) value(?, ?, ?, ?) ", post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}
