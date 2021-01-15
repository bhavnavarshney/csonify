package jsonio

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteJSONFile : Writes content in JSON File
func WriteJSONFile(csvFilePath string, writerChannel <-chan map[string]string, done chan<- bool, pretty bool) {
	writeString := createStringWriter(csvFilePath)
	jsonFunc, breakline := getJSONFunc(pretty)
	writeString("["+breakline, false)
	first := true

	for {
		record, more := <-writerChannel
		if more {
			if !first {
				writeString(","+breakline, false)
			} else {
				first = false
			}
			jsonData := jsonFunc(record)
			writeString(jsonData, false)
		} else {
			writeString(breakline+"]", false)
			done <- true
			break
		}
	}
}

func createStringWriter(csvPath string) func(string, bool) {
	jsonDir := filepath.Dir(csvPath)
	jsonName := fmt.Sprintf("%s.json", strings.TrimSuffix(filepath.Base(csvPath), ".csv"))
	final := filepath.Join(jsonDir, jsonName)
	file, err := os.Create(final)
	if err != nil {
		fmt.Println("Error creating JSON File")
		os.Exit(1)
	}
	return func(data string, isClose bool) {
		_, err := file.WriteString(data)
		if err != nil {
			fmt.Println("Internal Error")
			os.Exit(1)
		}
		if isClose {
			file.Close()
		}
	}

}

func getJSONFunc(pretty bool) (func(map[string]string) string, string) {
	var jsonFunc func(map[string]string) string
	var breakLine string
	if pretty {
		breakLine = "\n"
		jsonFunc = func(record map[string]string) string {
			jsonData, _ := json.MarshalIndent(record, "   ", "   ")
			return "   " + string(jsonData)
		}
	} else {
		breakLine = ""
		jsonFunc = func(record map[string]string) string {
			jsonData, _ := json.Marshal(record)
			return string(jsonData)
		}
	}
	return jsonFunc, breakLine
}
