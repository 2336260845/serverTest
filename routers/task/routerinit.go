package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"serverTest/apis"
	"serverTest/functions"
)

func TaskRouteInit(engine *gin.Engine) {
	engine.POST("/setTask", SetDelayEmailTask)
	engine.GET("/email", EmailSet)
}

func EmailSet(ctx *gin.Context) {
	s := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h1>输入待发送的邮件</h1>
发送邮件至:<input type="text", id="receivers"></br>
多少秒之后发送:<input type="text", id="delayTime"></br>
邮件标题:<input type="text", id="emailTitle"></br>
邮件内容:<input type="text", id="emailBody"></br>
<button id="sendInfobtn" onclick="SetSendEmailTask()" >设置任务</button>
<script>
    localPath = "http://127.0.0.1:3333";

    function SetSendEmailTask() {
        receviers = document.getElementById("receivers").value;
        delayTime = Number(document.getElementById("delayTime").value);
        emailTitle = document.getElementById("emailTitle").value;
        emailBody = document.getElementById("emailBody").value;

        var url = localPath + "/setTask";

        var orderData = {
            "title": emailTitle,
            "body" : emailBody,
            "receivers": receviers,
            "delayTime": delayTime
        };

        var jsonStr = JSON.stringify(orderData);

        xhr = new XMLHttpRequest();
        xhr.open("post", url, true);
        xhr.setRequestHeader('Content-Type', 'application/json');

        xhr.send(jsonStr)
    }
</script>
</body>
</html>

`
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, s)
	return
}

func SetDelayEmailTask(ctx *gin.Context) {
	var sit SimpleTaskStruct

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Errorf("读取body内容失败,err=%+v, body=%+v", err.Error(), string(body))
		ctx.JSON(500, apis.SendJson(errors.New("读取请求body失败"), err))
		return
	}

	err = json.Unmarshal(body, &sit)
	if err != nil {
		log.Errorf("Unmarshal失败,err=%+v, body=%+v", err.Error(), string(body))
		ctx.JSON(500, apis.SendJson(errors.New("Unmarshal失败,请校验请求参数是否正确"), err))
		return
	}

	log.Infof("sit=%+v", sit)
	if sit.Body == "" || sit.Receivers == "" || sit.Title == "" {
		log.Errorf("传入参数不正确,邮件标题,内容,接受人不能为空")
		ctx.JSON(400, apis.SendJson(errors.New("传入参数不正确,邮件标题,内容,接受人不能为空"), nil))
		return
	}

	eb := functions.NewDefaultEmailBody(sit.Title, sit.Body, []string{sit.Receivers})
	go functions.SetDelaySendEmail(eb, sit.DelayTime)

	ctx.JSON(200, apis.SendJson(nil, fmt.Sprintf("任务设置成功,将于%+vs后发送邮件", sit.DelayTime)))
}
