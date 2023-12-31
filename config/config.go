package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Setting struct {
	ListenDirectory         string `json:"listen_directory"`
	MarginFileNamePrefix    string `json:"margin_filename_prefix"`
	ShortSellFileNamePrefix string `json:"shortsell_filename_prefix"`
	MarginTableName         string `json:"margin_table_name"`
	ShortSellTableName      string `json:"shortsell_table_name"`
}

type Configuration struct {
	Settings *Setting `json:"settings"`
}

func BuildSettings() *Setting {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	stream, err := os.ReadFile(filepath.Join(basePath, "config", "config.json"))
	if err != nil {
		panic(err)
	}

	config := new(Configuration)

	err = json.Unmarshal(stream, &config)
	if err != nil {
		panic(err)
	}

	return config.Settings
}
