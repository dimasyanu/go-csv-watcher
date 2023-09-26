package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dimasyanu/go-csv-watcher/config"
	"github.com/dimasyanu/go-csv-watcher/contracts"
	"github.com/dimasyanu/go-csv-watcher/models"
	"github.com/dimasyanu/go-csv-watcher/modules"
	"github.com/dimasyanu/go-csv-watcher/utils"
)

var _settings = config.BuildSettings()
var _helper = utils.CreateHelper(_settings)
var _dirWatcher = modules.CreateDirWatcherModule(_settings, &_helper)
var _csvEngine = modules.CreateCsvReaderModule(_settings, &_helper)

func main() {
	Setup()
	WatchDirectory()
}

func Setup() {
	fmt.Println("Service started at:", time.Now().Format(time.RFC3339))
}

func WatchDirectory() {
	fileNamePrefixes := []string{_settings.MarginFileNamePrefix, _settings.ShortSellFileNamePrefix}
	files := _dirWatcher.GetFilesByPrefixes(_helper.GetListenDir(), fileNamePrefixes)
	if len(files) < 1 {
		return
	}

	for _, file := range files {
		items := _csvEngine.ExtractCsvData(file, func(fileName string) contracts.ICsvModel {
			if strings.HasPrefix(fileName, _settings.MarginFileNamePrefix) {
				return models.CreateMarginCsv()
			} else if strings.HasPrefix(fileName, _settings.ShortSellFileNamePrefix) {
				return models.CreateShortSellCsv()
			} else {
				panic("File model not found: " + fileName)
			}
		})
		if items == nil {
		}
		for _, item := range items {
			fmt.Println(item)
		}
	}
}
