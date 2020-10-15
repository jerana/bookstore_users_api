package services

import (
	"github.com/jerana/bookstore_users_api/domain/users"
	"github.com/jerana/bookstore_users_api/utils/errors"
)

//CreateUser create user in db
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := users.Validate(&user); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

//GetUser get user Info
func GetUser(userId int64) (*users.User, *errors.RestErr) {

	if userId <= 0 {
		return nil, errors.NewBadRequestError("Invalid User Id")
	}
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
