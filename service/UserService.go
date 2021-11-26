package service

import (
	"mark-v1/common/db"
	"mark-v1/models"
)

type UserService struct {
}

func (UserService UserService) GetUserByName(username string) models.Users {

	var user models.Users

	sqlStr := "select user_name,password,age,id,address,phone from users where user_name=?"

	error := db.Db.QueryRow(sqlStr, username).Scan(&user.UserName, &user.Password, &user.Age, &user.Id, &user.Address, &user.Phone)

	if error != nil {

	}

	return user
}