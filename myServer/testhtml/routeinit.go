package testhtml

import "github.com/gin-gonic/gin"

func TestRouteInit(engine *gin.Engine) {
	engine.Any("/", Hello)
	engine.Any("/button", Hello2)
}

