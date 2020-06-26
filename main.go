package main

import (
	"github.com/BohdanPatrash/indenticon/handler"
	"github.com/BohdanPatrash/indenticon/repos"
	"github.com/dgraph-io/badger"
	"github.com/gin-gonic/gin"
)

var db *badger.DB

func main() {
	db = repos.Init()
	defer db.Close()

	r := gin.Default()
	r.GET("/", handler.HomeGet())
	r.POST("/", handler.HomePost())
	r.Run()
}
