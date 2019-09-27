package file

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"serverTest/apis"
	"serverTest/conf"
)

func FileOpRouteInit(engine *gin.Engine) {
	engine.GET("/home/fileopt", Fileopthtml)
	engine.POST("/home/fileuplaod", Fileupload)
	engine.LoadHTMLGlob("view/*")
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


	err := ctx.Request.ParseMultipartForm(1 << 32)
	if err != nil {
		log.Errorf("ParseMultipartForm err:%+v", err.Error())
	}

	if ctx.Request.MultipartForm == nil {
		log.Errorf("MultipartForm is empty")
		ctx.JSON(400, apis.SendJson(errors.New("没有找到文件"), nil))
		return
	}
	fileList := ctx.Request.MultipartForm.File["file"]

	if len(fileList) <= 0 {
		log.Errorf("没有找到file文件,fileList=%+v", fileList)
		ctx.JSON(400, apis.SendJson(errors.New("没有找到file文件"), fileList))
		return
	}

	if len(fileList) != 1 {
		log.Errorf("不支持多个file一起上传,fileList=%+v", fileList)
		ctx.JSON(400, apis.SendJson(errors.New("不支持多个file一起上传"), fileList))
		return
	}

	fileRec := ctx.Request.MultipartForm.File["file"][0]

	fileType := fileRec.Header.Get("Content-Type")
	if fileType == "" {
		log.Errorf("无法得到文件格式,fileType=%+v", fileType)
		ctx.JSON(400, apis.SendJson(errors.New("无法得到文件格式,header头部信息参考data"), fileRec.Header))
		return
	}

	if fileType == "image/jpeg" {
		fileop, err := fileRec.Open()
		if err != nil {
			log.Errorf("无法打开文件,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法打开文件,报错信息参考data"), err.Error()))
			return
		}

		image, err := jpeg.Decode(fileop)
		if err != nil {
			log.Errorf("无法解析文件,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法解析文件,报错信息参考data"), err.Error()))
			return
		}

		cfg := conf.Cfg
		fileNew, err := os.Create(cfg.ProjectPath+"/static/tmp.jpg")
		if err != nil {
			log.Errorf("无法新建临时文件,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法新建临时文件,报错信息参考data"), err.Error()))
			return
		}

		err = jpeg.Encode(fileNew, image, nil)
		if err != nil {
			log.Errorf("无法保存图片,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法保存图片,报错信息参考data"), err.Error()))
			return
		}
	}

	log.Infof("保存图片成功")
	ctx.JSON(200, apis.SendJson(nil, nil))
	return
}

func FilePull(ctx *gin.Context) {

}
