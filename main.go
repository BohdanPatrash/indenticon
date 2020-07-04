package main

import (
	"github.com/BohdanPatrash/indenticon/handler"
	"github.com/BohdanPatrash/indenticon/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	db := repo.Init()
	defer db.Close()

	repo.DatabaseSetup()

	r := gin.Default()
	r.GET("/user", handler.UserGet())
	r.POST("/user", handler.UserPost())
	r.Run()
}
