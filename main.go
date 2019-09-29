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

	_, err := toml.DecodeFile(cfgFile, &conf.Cfg)
	if err != nil {
		fmt.Printf("toml.DecodeFile: %+v\n", err)
		return
	}

	server := routers.NewGinEngine()
	if err := server.Run(fmt.Sprintf("%s:%d", "0.0.0.0", conf.Cfg.ListenAddr)); err != nil {
		log.Errorf("服务启动失败")
		os.Exit(-1)
	}
}
