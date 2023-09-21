package utils

import (
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/dimasyanu/go-csv-watcher/config"
)

type DirWatcherProcessEngine struct {
	Settings *config.Setting
	BasePath string
}

func CreateDirWatcher(setting *config.Setting) DirWatcherProcessEngine {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return DirWatcherProcessEngine{
		Settings: setting,
		BasePath: basePath,
	}
}

func (o DirWatcherProcessEngine) FileExists() bool {
	return false
}

func (o DirWatcherProcessEngine) GetFiles() []fs.FileInfo {
	basePath := o.BasePath
	listenDir := o.Settings.ListenDirectory

	// Additional slash
	if basePath[len(basePath)-1:] != "/" && listenDir[:1] != "/" {
		listenDir = "/" + listenDir
	}
	listenDir = basePath + listenDir

	// Create directory if it is not exists
	if _, err := os.Stat(listenDir); os.IsNotExist(err) {
		err := os.Mkdir(listenDir, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}

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
