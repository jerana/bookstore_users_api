package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerana/bookstore_users_api/domain/users"
	"github.com/jerana/bookstore_users_api/services"
	"github.com/jerana/bookstore_users_api/utils/errors"
)

//CreateUser create user
func CreateUser(c *gin.Context) {
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	result, serErr := services.CreateUser(u)
	if serErr != nil {
		c.JSON(serErr.Status, serErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser get user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
