package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bhavnavarshney/csonify/csv"
	"github.com/bhavnavarshney/csonify/jsonio"
)

func conversion() {
	// For --options
	if len(os.Args) < 2 {
		fmt.Println("Please provide argument in format ./csonify [options] <csv-file>")
		return
	}
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	fileData, err := csv.GetFileData()
	if err != nil {
		os.Exit(1)
	}
	if _, err := csv.CheckIfValidFile(fileData.FilePath); err != nil {
		os.Exit(1)
	}
	writerChannel := make(chan map[string]string)
	done := make(chan bool)
	go csv.Process(fileData, writerChannel)
	go jsonio.WriteJSONFile(fileData.FilePath, writerChannel, done, fileData.Pretty)
	<-done
}
