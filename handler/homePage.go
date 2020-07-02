package handler

import (
	"fmt"

	"github.com/BohdanPatrash/indenticon/service"
	"github.com/gin-gonic/gin"
)

//HomeGet is handler function for GET request on "/"
func HomeGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		// users := service.GetAll()
		// fmt.Println(users)
		// c.JSON(http.StatusOK, users)
		service.CreateImage([]byte{})
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
		service.AddUser(email.Address)
		err := service.CreateImage(service.GetAll()[0].Hash)
		if err != nil {
			fmt.Println(err)
		}
	}
}
