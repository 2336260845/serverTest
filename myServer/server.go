package myServer

import (
	"github.com/gin-gonic/gin"
	"serverTest/myServer/file"
	"serverTest/myServer/testhtml"
)

func NewGinEngine() (engine *gin.Engine) {
	server := gin.Default()
	createRouter(server)
	return server
}

func createRouter(engine *gin.Engine) {
	//第一个测试服务
	testhtml.TestRouteInit(engine)
	file.FileOpRouteInit(engine)
	return
}
