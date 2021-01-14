package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/bhavnavarshney/csonify/model"
)

// Process : Process CSV File
func Process(inputfile model.InputFile, writerChannel chan<- map[string]string) {
	file, err := os.Open(inputfile.FilePath)
	if err != nil {
		fmt.Printf("Unable to open file: %v", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	if inputfile.Separator == "semicolon" {
		csvReader.Comma = ';'
	}

	headers, err := csvReader.Read()
	if err != nil {
		fmt.Println("Unable to read CSV File")
		return
	}

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			close(writerChannel)
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		record, err := processLine(headers, line)
		if err != nil {
			fmt.Printf("Line: %s, Error:%v", line, err)
			continue
		}
		writerChannel <- record
	}

}

func processLine(headers []string, line []string) (map[string]string, error) {
	if len(headers) != len(line) {
		return nil, errors.New("line does not match the headers format. skipping")
	}
	recordMap := make(map[string]string)
	for index, value := range headers {
		recordMap[value] = line[index]
	}
	return recordMap, nil
}
