package main

import (
	"go-share-books/configs"
	"math/rand"
	"time"

	"go-share-books/web"
	"github.com/gin-gonic/gin"
)
func init() {
	rand.Seed(time.Now().UnixNano())
	config.Parse()
	// web.ThirdCli()
	// mission.AutoMigrate()
}
func main () {
	r := gin.New()

	web.Run(r)
}