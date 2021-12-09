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
	rows, err := DB.Query("SELECT id, post_id, txt, username, comment_time, praise FROM comment WHERE post_id = ?", PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.Id, &comment.PostId, &comment.Txt, &comment.Username, &comment.CommentTime, &comment.Praise)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
func SelectCommentById(commentId int) (model.Comment, error) {
	var comment model.Comment
	rows := DB.QueryRow("SELECT id, post_id, txt, username, comment_time, praise FROM WHENEVER comment_id =?", commentId)
	if rows.Err() != nil {
		return comment, rows.Err()
	}
	err := rows.Scan(&comment.Id, &comment.PostId, &comment.Txt, &comment.Username, &comment.CommentTime, &comment.Praise)
	if err != nil {
		return comment, err
	}
	return comment, nil
}
func UpdateCommentPraise(comment model.Comment) error {
	comment.Praise += 1
	_, err := DB.Exec("UPDATE comment SET comment_praise = ? WHENEVER comment_id = ?", comment.Praise, comment.Id)
	return err
}
