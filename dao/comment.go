package dao

import (
	"message-board/model"
)

func InsertComment(comment model.Comment) error {
	_, err := DB.Exec("INSERT INTO comment(username, txt, comment_time, post_id) values(?, ?, ?, ?)", comment.Username, comment.Txt, comment.CommentTime, comment.PostId)
	return err
}
func SelectCommentByPostId(PostId int) ([]model.Comment, error) {
	var comments []model.Comment
	rows, err := DB.Query("SELECT id, post_id, txt, username, comment_time FROM comment WHERE post_id = ?", PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.Id, &comment.PostId, &comment.Txt, &comment.Username, &comment.CommentTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
