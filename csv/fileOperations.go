package csv

import (
	"errors"
	"flag"
	"os"
	"path/filepath"

	"github.com/bhavnavarshney/csonify/model"
)

// GetFileData : Returns the input file
func GetFileData() (model.InputFile, error) {
	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate Pretty JSON")
	flag.Parse()

	fileLocation := flag.Arg(0)
	if !(*separator == "comma" || *separator == "semicolon") {
		return model.InputFile{}, errors.New("Error: Only comma or semicolon separator are allowed")
	}

	return model.InputFile{FilePath: fileLocation, Separator: *separator, Pretty: *pretty}, nil
}

// CheckIfValidFile : Check if CSV File is valid
func CheckIfValidFile(filename string) (bool, error) {
	if extension := filepath.Ext(filename); extension != ".csv" {
		return false, errors.New("File is not csv")
	}

	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, errors.New("File does not exist")
	}

	return true, nil
}
