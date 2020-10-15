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
