package dao

import (
	"message-board/model"
)

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("SELECT id, password FROM user WHERE username = ? ", username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
func InsertUser(user model.User) error {
	_, err := DB.Exec("INSERT INTO user(name,password) value (?,?)", user.Name, user.Password)
	return err
}
func UpdatePassword(username, password string) error {
	_, err := DB.Exec("UPDATE user SET password = ? WHENEVER username = ?", password, username)
	return err
}
