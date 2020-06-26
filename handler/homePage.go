package handler

import (
	"fmt"
	"net/http"

	"github.com/BohdanPatrash/indenticon/repos/indenticons"
	"github.com/gin-gonic/gin"
)

//HomeGet is handler function for GET request on "/"
func HomeGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		icon, err := indenticons.GetByEmail("some@email")
		if err != nil {
			fmt.Println("FAILED: ", err.Error())
			c.Status(http.StatusBadRequest)
		}
		c.JSON(http.StatusOK, icon)
	}
}

//HomePost is handler function for POST request on "/"
func HomePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Bind(&db)
	}
}
