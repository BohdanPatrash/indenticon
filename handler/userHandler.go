package handler

import (
	"net/http"

	"github.com/BohdanPatrash/indenticon/repo"
	"github.com/BohdanPatrash/indenticon/service"
	"github.com/gin-gonic/gin"
)

//UserGet is handler function for GET request on "/"
func UserGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := repo.GetAllUsers()
		c.JSON(http.StatusOK, users)
	}
}

//UserPost is handler function for POST request on "/"
func UserPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Address struct {
			Address string `json:"email"`
		}
		email := Address{}
		c.Bind(&email)
		user := service.CreateUser(email.Address)
		err := repo.SaveUser(user)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}
