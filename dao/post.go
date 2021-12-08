package dao

import "message-board/model"

func AddPost(post model.Post) error {
	_, err := DB.Exec("INSERT INTO post(username, txt, post_time, update_time) value(?, ?, ?, ?) ", post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}
func SelectPost() ([]model.Post, error) {
	var posts []model.Post
	rows, err := DB.Query("SELECT id, username, txt, post_time, update_time, comment_num FROM post ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post model.Post
		err = rows.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func SelectPostById(postId int) (model.Post, error) {
	var post model.Post
	row := DB.QueryRow("SELECT id, username, txt, post_time, update_time, comment_num FROM post WHERE id = ? ", postId)
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum)
	if err != nil {
		return post, err
	}
	return post, nil
}
