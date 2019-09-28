package file

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"serverTest/apis"
	"serverTest/conf"
	"serverTest/functions"
	"strconv"
	"time"
)

func FileOpRouteInit(engine *gin.Engine) {
	engine.GET("/home/fileopt", Fileopthtml)
	//engine.LoadHTMLGlob("view/*")
	engine.POST("/fileop/push", FilePush)
	engine.POST("/fileop/pull", FilePull)
	engine.StaticFS("/upload/images", http.Dir(GetImageFullPath()))
}

func GetImageFullPath() string {
	cfg := conf.Cfg
	path := cfg.ProjectPath + "/static/"
	log.Infof("GetImageFullPath=%+v", path)
	return path
}

func Fileopthtml(c *gin.Context) {
	c.HTML(http.StatusOK, "fileopt.html", gin.H{
		"title": "GIN: 文件上传下载操作布局页面",
	})
}

func GetDateString() string {
	currentTime := time.Now().Unix()
	return strconv.FormatInt(currentTime, 10)
}

func FilePush(ctx *gin.Context) {
	log.Infof("开始上传图片")

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
		log.Infof("图片为image/jpeg格式,开始保存")
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

		//保存到挂载目录中
		fileUpload, err := os.Create(cfg.ProjectPath + "/static/uploadfile/" + fileRec.Filename)
		if err != nil {
			log.Errorf("无法新建临时文件,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法新建临时文件,报错信息参考data"), err.Error()))
			return
		}

		err = jpeg.Encode(fileUpload, image, nil)
		if err != nil {
			log.Errorf("无法保存图片,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法保存图片,报错信息参考data"), err.Error()))
			return
		}

		//保存到history中
		fileNew, err := os.Create(cfg.ProjectPath + "/static/historyfile/" + GetDateString() + ".jpg")
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

		log.Infof("保存图片成功")

	} else if fileType == "image/png" {
		fileop, err := fileRec.Open()
		if err != nil {
			log.Errorf("无法打开文件,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法打开文件,报错信息参考data"), err.Error()))
			return
		}

		image, err := png.Decode(fileop)
		if err != nil {
			log.Errorf("无法解析文件,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法解析文件,报错信息参考data"), err.Error()))
			return
		}

		cfg := conf.Cfg

		//保存到挂载目录中
		fileUpload, err := os.Create(cfg.ProjectPath + "/static/uploadfile/" + fileRec.Filename)
		if err != nil {
			log.Errorf("无法新建临时文件,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法新建临时文件,报错信息参考data"), err.Error()))
			return
		}

		err = png.Encode(fileUpload, image)
		if err != nil {
			log.Errorf("无法保存图片,err=%+v", err.Error())
			ctx.JSON(400, apis.SendJson(errors.New("无法保存图片,报错信息参考data"), err.Error()))
			return
		}

		//保存到history中
		fileNew, err := os.Create(cfg.ProjectPath + "/static/historyfile/" + GetDateString() + ".png")
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

		log.Infof("保存图片成功")
	} else {
		log.Infof("不支持的图片格式:%+v", fileType)
		ctx.JSON(400, apis.SendJson(fmt.Errorf("不支持的图片格式:%+v", fileType), nil))
		return
	}

	ctx.JSON(200, apis.SendJson(nil, nil))
	return
}

func FilePull(ctx *gin.Context) {
	log.Infof("开始下载图片")

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
		log.Errorf("不支持多个file一起下载,fileList=%+v", fileList)
		ctx.JSON(400, apis.SendJson(errors.New("不支持多个file一起下载"), fileList))
		return
	}

	fileRec := ctx.Request.MultipartForm.File["file"][0]

	fileType := fileRec.Header.Get("Content-Type")
	if fileType == "" {
		log.Errorf("无法得到文件格式,fileType=%+v", fileType)
		ctx.JSON(400, apis.SendJson(errors.New("无法得到文件格式,header头部信息参考data"), fileRec.Header))
		return
	}

	err = functions.ImageCrop(fileRec.Filename)
	if err != nil {
		log.Errorf("自动切割图片失败,fileName:%+v,err=%+v", fileRec.Filename, err.Error())
		ctx.JSON(500, apis.SendJson(fmt.Errorf("自动切割图片失败,fileName:%+v,err=%+v", fileRec.Filename, err.Error()), nil))
		return
	}

	log.Infof("完成自动切割图片:%+v", fileRec.Filename)
	ctx.JSON(200, apis.SendJson(nil, nil))
}
