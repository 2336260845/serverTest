package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func FileOpRouteInit(engine *gin.Engine) {
	engine.GET("/home/fileopt", Fileopthtml)
	engine.POST("/home/fileuplaod", Fileupload)
	engine.LoadHTMLGlob("../view/*")
	engine.POST("/fileop/push", FilePush)
	engine.POST("/fileop/pull")
}

func Fileopthtml(c *gin.Context) {
	c.HTML(http.StatusOK, "fileopt.html", gin.H{
		"title": "GIN: 文件上传下载操作布局页面",
	})
}

func Fileupload(c *gin.Context) {
	//得到上传的文件
	file, header, err := c.Request.FormFile("image") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename

	fmt.Println(file, err, filename)
	//创建文件
	out, err := os.Create("static/uploadfile/" + filename)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusCreated, "upload successful")
}

func FilePush(ctx *gin.Context) {
	funcName := "FilePush"

	fileType := ctx.Request.Header.Get("Content-Type")
	header := ctx.Request.Header

	if fileType == "image/jpeg" {
		log.Infof("图片格式为:", fileType)
		img, str, err := image.Decode(ctx.Request.Body)
		if err != nil {
			log.Errorf("%s:decode image error,err=%+v", funcName, err.Error())
		}

		log.Infof("str=%+v, img=%+v", str, img)
	} else {
		log.Infof("不支持的图片类型,fileType=%+v, header=%+v\n\n", fileType, header)
		log.Infof("param=%+v\n\n", ctx.Params)

		body, _ := ioutil.ReadAll(ctx.Request.Body)
		log.Infof("body=%+v\n\n", string(body))

	}

}

func FilePull(ctx *gin.Context) {

}
