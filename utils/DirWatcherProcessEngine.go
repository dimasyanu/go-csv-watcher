package utils

import (
	"io/fs"
	"io/ioutil"

	"github.com/dimasyanu/go-csv-watcher/config"
)

type DirWatcherProcessEngine struct {
	Settings *config.Setting
	Helper   *Helper
}

func CreateDirWatcher(setting *config.Setting, helper *Helper) DirWatcherProcessEngine {
	return DirWatcherProcessEngine{
		Settings: setting,
		Helper:   helper,
	}
}

func (o DirWatcherProcessEngine) FileExists() bool {
	return false
}

func (o DirWatcherProcessEngine) GetFiles() []fs.FileInfo {
	listenDir := o.Helper.GetListenDir()

	// List files
	files, err := ioutil.ReadDir(listenDir)
	if err != nil {
		panic(err)
	}

	results := []fs.FileInfo{}

	// Make result list
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		results = append(results, file)
	}

	return results
}
