package utils

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
	"github.com/dimasyanu/go-csv-watcher/models"
)

type CsvProcessEngine struct {
	Settings *config.Setting
	Helper   *Helper
}

func CreateCsvProcessEngine(settings *config.Setting, helper *Helper) *CsvProcessEngine {
	return &CsvProcessEngine{
		Settings: settings,
		Helper:   helper,
	}
}

func (o CsvProcessEngine) ExtractCsvData(file fs.FileInfo) []models.ICsvModel {
	fileStream, err := os.Open(o.Helper.GetListenDir() + "/" + file.Name())
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(fileStream)
	fileScanner.Split(bufio.ScanLines)

	results := []models.ICsvModel{}
	i := 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if i > 10 || strings.Contains(text, "Efek;board") {
			continue
		}
		item := o.ParseTextIntoModel(file.Name(), text)
		results = append(results, item)
		i++
	}

	fileStream.Close()

	return results
}

func (o CsvProcessEngine) ParseTextIntoModel(fileName string, text string) models.ICsvModel {
	// Turn each line into array of values
	segments := strings.Split(text, ";")
	if false {
		fmt.Println(segments)
	}

	// Determine model type
	var newItem models.ICsvModel
	if strings.HasPrefix(fileName, o.Settings.MarginFileNamePrefix) {
		newItem = models.CreateMarginCsv()
	}

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
