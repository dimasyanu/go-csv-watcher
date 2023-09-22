package utils

import (
	"os"

	"github.com/dimasyanu/go-csv-watcher/config"
)

type Helper struct {
	Settings *config.Setting
	BasePath string
}

func CreateHelper(settings *config.Setting) Helper {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return Helper{
		Settings: settings,
		BasePath: basePath,
	}
}

func (o Helper) GetListenDir() string {
	basePath := o.BasePath
	listenDir := o.Settings.ListenDirectory

	// Additional slash
	if basePath[len(basePath)-1:] != "/" && listenDir[:1] != "/" {
		listenDir = "/" + listenDir
	}
	listenDir = basePath + listenDir

	// Create directory if it is not exists
	if _, err := os.Stat(listenDir); os.IsNotExist(err) {
		err := os.Mkdir(listenDir, 0775)
		if err != nil {
			panic(err)
		}
	}

	return listenDir
}
