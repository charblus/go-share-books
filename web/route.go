package web

import (
	"go-share-books/web/controller/origin"
	// "go-share-books/web/controller/index"
	// "go-share-books/web/controller/login"

	"github.com/gin-gonic/gin"
)


//Route 映射路由
func Route(engine *gin.Engine) {
	r := engine.Group("/api") 
	{
		r.GET("/origin", origin.GetOrigin)
		r.POST("/index", origin.GetOrigin)
	}
}
