package modules

import (
	"io/fs"
	"os"
	"strings"

	"github.com/dimasyanu/go-csv-watcher/config"
	"github.com/dimasyanu/go-csv-watcher/utils"
)

type DirWatcherModule struct {
	Settings *config.Setting
	Helper   *utils.Helper
}

func CreateDirWatcherModule(setting *config.Setting, helper *utils.Helper) *DirWatcherModule {
	return &DirWatcherModule{
		Settings: setting,
		Helper:   helper,
	}
}

func (o DirWatcherModule) FileExists() bool {
	return false
}

func (o DirWatcherModule) GetFiles(listenDir string) []fs.DirEntry {

	// List files
	files, err := os.ReadDir(listenDir)
	if err != nil {
		panic(err)
	}

	results := []fs.DirEntry{}

	// Make result list
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		results = append(results, file)
	}

	return results
}

func (o DirWatcherModule) GetFilesByPrefix(listenDir string, prefix string) []fs.DirEntry {
	files := o.GetFiles(listenDir)
	results := []fs.DirEntry{}
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), prefix) {
			continue
		}
		results = append(results, file)
	}

	return results
}

func (o DirWatcherModule) GetFilesByPrefixes(listenDir string, prefixes []string) []fs.DirEntry {
	results := []fs.DirEntry{}

	files := o.GetFiles(listenDir)
	for _, file := range files {
		processFile := false
		for _, namePrefix := range prefixes {
			if strings.HasPrefix(file.Name(), namePrefix) {
				processFile = true
				break
			}
		}
		if processFile {
			results = append(results, file)
		}
	}
	return results
}
