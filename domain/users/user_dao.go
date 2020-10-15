package users

import (
	"fmt"

	"github.com/jerana/bookstore_users_api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

//GET get user
func (u *User) Get() *errors.RestErr {
	result := userDB[u.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}
	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	return nil
}

//Save save user
func (u *User) Save() *errors.RestErr {
	current := userDB[u.Id]
	if current != nil {
		if current.Email != u.Email {
			return errors.NewBadRequestError(fmt.Sprint("Invalid Email Id %d for user", u.Id))
		}
		return errors.NewBadRequestError(fmt.Sprint("User %d already exit", u.Id))
	}

	userDB[u.Id] = u
	return nil
}
