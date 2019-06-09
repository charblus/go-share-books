package web

import (
	"go-share-books/configs"
	"go-share-books/db"
	// "go-share-books/web/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

//ThirdCli 第三方客户端
func ThirdCli() {
	if err := db.InitPsqlClient(config.Conf.Psql); err != nil {
		log.Fatal(err)
	}

	if err := db.IntiRedisClient(config.Conf.Redis); err != nil {
		// log.Fatal(err)
	}

	// 开启DB日志模式
	db.DB.LogMode(!config.Conf.Release)

}

// Middle 注册中间组件
func Middle(r *gin.Engine) {
	// 生产环境下配置
	if config.Conf.Release {
		gin.SetMode(gin.ReleaseMode) // 设置框架环境
	} else {
		r.Use(gin.Logger())
	}
	// r.Use(middleware.RecoverPanic())

}

// Run 监听端口运行服务
func Run(r *gin.Engine) {
	addr := fmt.Sprintf(":%d", config.Conf.BindPort)
	mode := "development"
	if config.Conf.Release {
		mode = "production"
	}

	log.Printf("listening on %s (%s)", addr, mode)
	r.Run(addr)
}
