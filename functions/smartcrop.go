package functions

import (
	"fmt"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"image"
	"image/jpeg"
	log "github.com/sirupsen/logrus"
	"os"
	"serverTest/conf"
)

//自动完成图片切割
func ImageCrop(fileName string) error {
	funcName := "ImageCrop"
	cfg := conf.Cfg

	f, err := os.Open(cfg.ProjectPath+"/static/uploadfile/"+fileName)
	if err != nil {
		log.Errorf("%s:打开文件失败,err=%+v", funcName,err.Error())
		return fmt.Errorf("%s:打开文件失败,err=%+v", funcName,err.Error())
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Errorf("%s:解码文件失败,err=%+v", funcName,err.Error())
		return fmt.Errorf("%s:解码文件失败,err=%+v", funcName,err.Error())
	}

	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, _ := analyzer.FindBestCrop(img, 250, 250)

	log.Infof("Top crop: %+v\n", topCrop)

	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	croppedimg := img.(SubImager).SubImage(topCrop)
	f2, err := os.Create(cfg.ProjectPath+"/static/cropfiles/"+fileName)
	if err != nil {
		log.Errorf("create file err:%+v", err.Error())
		return fmt.Errorf("create file err:%+v", err.Error())
	}

	err = jpeg.Encode(f2, croppedimg, &jpeg.Options{100})
	if err != nil {
		log.Errorf("jpeg.Encode err:%+v", err.Error())
		return fmt.Errorf("jpeg.Encode err:%+v", err.Error())
	}

	return nil
}
