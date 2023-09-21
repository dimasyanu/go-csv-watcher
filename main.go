package main

import (
	"config"
	"fmt"
)

var _config config.Configuration

func main() {
	WatchDirectory()
}

func WatchDirectory() {
	fmt.Println(_config.Settings.ListenDirectory)
}
