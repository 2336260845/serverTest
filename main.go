package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"playServer/myServer"
	"os"
)

func main() {
	//cfg := conf.Config{}

	server := myServer.NewGinEngine()
	if err := server.Run(fmt.Sprintf("%s:%d", "0.0.0.0", 3333)); err != nil {
		log.Errorf("服务启动失败")
		os.Exit(-1)
	}
}
