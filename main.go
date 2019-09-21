package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"os"
	"serverTest/conf"
	"serverTest/routers"
)

var versionStr = "unknown"
var commitStr = "unknown"
var cfgFile = ""

func parseFlag() {
	ver := flag.Bool("v", false, "version")
	help := flag.Bool("h", false, "help")
	flag.StringVar(&commitStr, "commit", commitStr, "commit")
	flag.StringVar(&cfgFile, "config", "./config.toml", "configFile path")
	flag.Parse()

	if *ver {
		fmt.Println(versionStr)
		os.Exit(0)
	} else if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {
	parseFlag()
	cfg := conf.Config{}

	_, err := toml.DecodeFile(cfgFile, &cfg)
	if err != nil {
		fmt.Printf("toml.DecodeFile: %+v\n", err)
		return
	}

	//------------------------------------------------------------------------------------
/*
	f, _ := os.Open("123.jpg")
	img, _, _ := image.Decode(f)


	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, _ := analyzer.FindBestCrop(img, 250, 250)

	fmt.Printf("Top crop: %+v\n", topCrop)

	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	croppedimg := img.(SubImager).SubImage(topCrop)
	f2, err := os.Create("123_test.jpg")
	if err != nil {
		fmt.Printf("create file err:%+v", err.Error())
		os.Exit(1)
	}

	err = jpeg.Encode(f2, croppedimg, &jpeg.Options{100})
	if err != nil {
		fmt.Printf("jpeg.Encode err:%+v", err.Error())
		os.Exit(1)
	}
*/
	//------------------------------------------------------------------------------------

	server := routers.NewGinEngine()
	if err := server.Run(fmt.Sprintf("%s:%d", "0.0.0.0", cfg.ListenAddr)); err != nil {
		log.Errorf("服务启动失败")
		os.Exit(-1)
	}
}
