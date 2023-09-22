package utils

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"strconv"
	"strings"

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
	// Turn text into array
	segments := strings.Split(text, ";")
	if false {
		fmt.Println(segments)
	}

	// Determine model type
	var objValues reflect.Value
	var objType reflect.Type
	var newItem models.ICsvModel
	if strings.HasPrefix(fileName, o.Settings.MarginFileNamePrefix) {
		newItem = models.MarginCsv{}
	}

	objValues = reflect.ValueOf(&newItem).Elem()
	// objType = reflect.TypeOf(newItem)

	// Map acquired segments into model
	for i := 0; i < objValues.NumField(); i++ {
		fieldVals := objValues.Field(i)
		// fieldInfo := objType.Field(i)

		if !fieldVals.IsValid() || !fieldVals.CanSet() {
			continue
		}

		// Parse float64
		if fieldVals.Kind() == reflect.Float64 {
			val, err := strconv.ParseFloat(segments[fieldInfo.Index[0]], 64)
			if err != nil {
				panic(err)
			}
			fieldVals.SetFloat(val)
			continue
		}

		// Parse int32
		if fieldVals.Kind() == reflect.Int32 {
			val, err := strconv.ParseInt(segments[fieldInfo.Index[0]], 0, 32)
			if err != nil {
				panic(err)
			}
			fieldVals.SetInt(val)
			continue
		}
	}

	fmt.Println(newItem)
	return models.MarginCsv{}
}
