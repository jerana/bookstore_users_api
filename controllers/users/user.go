package users

import (
	"net/http"
	"strconv"

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
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user Id is not valid")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
