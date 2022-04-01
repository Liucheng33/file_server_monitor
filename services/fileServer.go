package services

import (
	"file_server_monitor/config"
	"file_server_monitor/utils"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	watcher    *fsnotify.Watcher
	fileFolder = "."
)

func WatchStart(cfg config.Config) {
	done := make(chan bool)
	//开始监听
	initWatcher(cfg)
	defer watcher.Close()
	<-done
	return
}

func initWatcher(cfg config.Config) {
	var err error
	if watcher != nil {
		_ = watcher.Close()
	}
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		utils.LogAndExit(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//go watchChangeHandler(event)
				eventDispatcher(event)
			}
		}
	}()
	addWatcher(cfg)

}

func eventDispatcher(event fsnotify.Event) {
	//todo  发送消息推送MQ
	log.Println("发送消息给MQ")
	log.Println(event.Name)

}

/*
func watchChangeHandler(event fsnotify.Event) {
	if event.Op != fsnotify.Create && event.Op != fsnotify.Rename {
		return
	}
	_, err := ioutil.ReadDir(event.Name)
	if err != nil {
		return
	}

}

*/

func addWatcher(cfg config.Config) {
	log.Println("collecting directory information...")
	var err error
	fileFolder, err = os.Getwd()
	if err != nil {
		utils.LogAndExit(err)
	}
	dirsMap := map[string]bool{}
	for _, dir := range cfg.Monitor.IncludeDirs {
		darr := utils.DirParse2Array(dir)
		if len(darr) < 1 || len(darr) > 2 {
			utils.LogAndExit("config section monitor dirs is error. ", dir)
		}
		if strings.HasPrefix(darr[0], "/") {
			utils.LogAndExit("dirs must be relative paths ! err path:", dir)
		}
		if darr[0] == "." {
			if len(darr) == 2 && darr[1] == "*" {
				// The highest priority
				dirsMap = map[string]bool{
					fileFolder: true,
				}
				utils.ListFile(fileFolder, func(d string) {
					dirsMap[d] = true
				})
				cfg.Monitor.IncludeDirsRec = map[string]bool{}
				cfg.Monitor.IncludeDirsRec[fileFolder] = true
				break
			} else {
				dirsMap[fileFolder] = true
			}
		} else {
			md := fileFolder + "\\" + darr[0]
			dirsMap[md] = true
			if len(darr) == 2 && darr[1] == "*" {
				utils.ListFile(md, func(d string) {
					dirsMap[d] = true
				})
				cfg.Monitor.IncludeDirsRec = map[string]bool{}
				cfg.Monitor.IncludeDirsRec[fileFolder] = true
			}
		}

	}

	for dir := range dirsMap {
		log.Println("watcher add -> ", dir)
		err := watcher.Add(dir)
		if err != nil {
			utils.LogAndExit(err)
		}
	}
	log.Println("total monitored dirs: " + strconv.Itoa(len(dirsMap)))
	log.Println("fsnotify is ready.")
	cfg.Monitor.DirsMap = dirsMap
}
