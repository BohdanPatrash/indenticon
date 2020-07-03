package handler

import (
	"net/http"

	"github.com/BohdanPatrash/indenticon/service"
	"github.com/gin-gonic/gin"
)

//HomeGet is handler function for GET request on "/"
func HomeGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := service.GetAll()
		c.JSON(http.StatusOK, users)
	}
}

//HomePost is handler function for POST request on "/"
func HomePost() gin.HandlerFunc {
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
