package modules

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dimasyanu/go-csv-watcher/config"
	"github.com/dimasyanu/go-csv-watcher/contracts"
	"github.com/dimasyanu/go-csv-watcher/utils"
)

type CsvReaderModule struct {
	Settings *config.Setting
	Helper   *utils.Helper
}

func CreateCsvReaderModule(settings *config.Setting, helper *utils.Helper) *CsvReaderModule {
	return &CsvReaderModule{
		Settings: settings,
		Helper:   helper,
	}
}

func (o CsvReaderModule) ExtractCsvData(file fs.DirEntry, fn func(string) contracts.ICsvModel) []contracts.ICsvModel {
	fileStream, err := os.Open(o.Helper.GetListenDir() + "/" + file.Name())
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	results := []contracts.ICsvModel{}
	i := 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if i > 10 || strings.Contains(text, "Efek;board") {
			continue
		}
		item := o.ParseTextIntoModel(file.Name(), text, fn)
		results = append(results, item)
		i++
	}

	fileStream.Close()

	return results
}

func (o CsvReaderModule) ParseTextIntoModel(fileName string, text string, getModelByFileName func(string) contracts.ICsvModel) contracts.ICsvModel {
	// Turn each line into array of values
	segments := strings.Split(text, ";")
	if false {
		fmt.Println(segments)
	}

	// Determine model type
	var newItem = getModelByFileName(fileName)

	objInfo := reflect.ValueOf(newItem)
	objValues := objInfo.Elem()

	// Map acquired segments into model
	for i := 0; i < objValues.NumField(); i++ {
		fieldVal := objValues.Field(i)

		if !fieldVal.IsValid() || !fieldVal.CanSet() {
			continue
		}

		val := segments[i]
		kind := fieldVal.Kind()

		// fieldName := reflect.Indirect(objInfo).Type().Field(i).Name
		// structName := " | " + kind.String()
		// fmt.Println(fieldName, "|", val, "|", kind, structName)

		// Parse bool
		if fieldVal.Kind() == reflect.Bool {
			fieldVal.Set(reflect.ValueOf(val == "Y"))
			continue
		}

		// Parse int32
		if fieldVal.Kind() == reflect.Int32 {
			val, err := strconv.ParseInt(segments[i], 0, 32)
			if err != nil {
				panic(err)
			}
			fieldVal.SetInt(val)
			continue
		}

		// Parse int64
		if fieldVal.Kind() == reflect.Int64 {
			val, err := strconv.ParseInt(segments[i], 0, 64)
			if err != nil {
				panic(err)
			}
			fieldVal.SetInt(val)
			continue
		}

		// Parse float64
		if fieldVal.Kind() == reflect.Float64 {
			val, err := strconv.ParseFloat(segments[i], 64)
			if err != nil {
				panic(err)
			}
			fieldVal.SetFloat(val)
			continue
		}

		// Parse time.Time
		if kind == reflect.Struct && (&fieldVal).Type() == reflect.TypeOf(time.Now()) {
			layout := "2006-01-02"
			if len(val) > 10 {
				layout += " 15:04:05"
			}
			timeVal, err := time.Parse(layout, val)
			if err != nil {
				panic(err)
			}
			fieldVal.Set(reflect.ValueOf(timeVal))
			continue
		}

		fieldVal.SetString(val)
	}

	// fmt.Println(newItem)
	return newItem
}
