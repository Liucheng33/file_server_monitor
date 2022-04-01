package main

import (
	"file_server_monitor/config"
	"file_server_monitor/services"
	"log"
	"os"
)

var (
	configPath = "etc/config-sample.yaml"
)

func main() {
	log.SetPrefix("file-fsnotify")
	log.SetFlags(2)
	log.SetOutput(os.Stdout)

	//获取配置
	var cfg config.Config
	config.LoadConfig(configPath, &cfg)
	//开启监听
	services.WatchStart(cfg)
}
