package handler

import (
	"net/http"

	"github.com/BohdanPatrash/indenticon/service"
	"github.com/gin-gonic/gin"
)

//UserGet is handler function for GET request on "/"
func UserGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := service.GetAll()
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
		service.SaveUser(user)
	}
}
