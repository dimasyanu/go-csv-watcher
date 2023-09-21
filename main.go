package main

import (
	"fmt"
	"time"

	"github.com/dimasyanu/go-csv-watcher/config"
	"github.com/dimasyanu/go-csv-watcher/utils"
)

var _settings = config.BuildSettings()
var _helper = utils.CreateHelper(_settings)
var _dirWatcher = utils.CreateDirWatcher(_settings)

func main() {
	Setup()
	WatchDirectory()
}

func Setup() {
	fmt.Println("Service started at:", time.Now().Format(time.RFC3339))
}

func WatchDirectory() {
	files := _dirWatcher.GetFiles()
	if len(files) < 1 {
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
