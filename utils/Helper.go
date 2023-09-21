package utils

import (
	"fmt"
	"time"

	"github.com/dimasyanu/go-csv-watcher/config"
)

type Helper struct {
	_settings *config.Setting
}

func CreateHelper(settings *config.Setting) Helper {
	return Helper{
		_settings: settings,
	}
}

func (o Helper) DateToString(datetime time.Time) string {
	now := time.Now()

	result := fmt.Sprint(
		"%d-%02d-%2dT%02d:%02d:%02d\n",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	)

	return result
}
