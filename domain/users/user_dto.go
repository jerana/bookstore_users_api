package users

import (
	"strings"

	"github.com/jerana/bookstore_users_api/utils/errors"
)

//User export used
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_Created"`
}

func Validate(user *User) *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	//fmt.Println("email: ", user.Email)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
