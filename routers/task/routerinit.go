package task

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"serverTest/apis"
	"serverTest/functions"
)

func TaskRouteInit(engine *gin.Engine) {
	engine.POST("/setTask", SetDelayEmailTask)
}

func SetDelayEmailTask(ctx *gin.Context) {
	var sit SimpleTaskStruct
	if err := ctx.ShouldBindJSON(&sit); err != nil {
		log.Errorf("请求参数不正确")
		ctx.JSON(400, apis.SendJson(errors.New("请求参数不正确"), nil))
		return
	}

	eb := functions.NewDefaultEmailBody(sit.Title, sit.Body, sit.Receivers)
	go functions.SetDelaySendEmail(eb, sit.DelayTime)

	ctx.JSON(200, apis.SendJson(nil, fmt.Sprintf("任务设置成功,将于%+vs后发送邮件", sit.DelayTime)))
}
