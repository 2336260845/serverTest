package testhtml

import "github.com/gin-gonic/gin"

func TestRouteInit(engine *gin.Engine) {
	engine.GET("/", Hello)
	engine.GET("/2336260845@qq.com", Writer)
	engine.GET("/file", File)
}
