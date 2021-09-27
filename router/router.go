package router

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	ginpprof.Wrap(r)
	r.GET("/ping", controller.Ping)

	return r
}
