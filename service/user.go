package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

//查询用户是否存在
func IsRepeatName(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func IsPasswordPlausible(user model.User) bool {
	if len(user.Password) < 8 {
		return false
	}
	return true
}
func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}
func IsPasswordRight(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}
func ChangePassword(username, password string) error {
	err := dao.UpdatePassword(username, password)
	return err
}
