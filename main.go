package main

import (
	"github.com/BohdanPatrash/indenticon/handler"
	"github.com/BohdanPatrash/indenticon/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	repo.Init()

	r := gin.Default()
	r.GET("/", handler.HomeGet())
	r.POST("/", handler.HomePost())
	r.Run()
}
