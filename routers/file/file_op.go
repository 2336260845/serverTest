package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func FileOpRouteInit(engine *gin.Engine) {
	engine.GET("/home/fileopt", Fileopthtml)
	engine.POST("/home/fileuplaod", Fileupload)
	engine.LoadHTMLGlob("view/*")
	engine.POST("/fileop/push", )
	engine.POST("/fileop/pull", )
}

func Fileopthtml(c *gin.Context){
	c.HTML(http.StatusOK, "fileopt.html", gin.H{
		"title": "GIN: 文件上传下载操作布局页面",
	})
}

func Fileupload(c *gin.Context){
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
	out, err := os.Create("static/uploadfile/"+filename)
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

}

func FilePull(ctx *gin.Context) {

}