package csv

import (
	"flag"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/bhavnavarshney/csonify/model"
)

func TestGetFileData(t *testing.T) {
	tests := []struct {
		name    string
		want    model.InputFile
		wantErr bool
		osArgs  []string
	}{
		{"Default parameters", model.InputFile{"test.csv", "comma", false}, false, []string{"cmd", "test.csv"}},
		{"Semicolon enabled", model.InputFile{"test.csv", "semicolon", false}, false, []string{"cmd", "--separator=semicolon", "test.csv"}},
		{"Pretty enabled", model.InputFile{"test.csv", "comma", true}, false, []string{"cmd", "--pretty", "test.csv"}},
		{"Pretty and semicolon enabled", model.InputFile{"test.csv", "semicolon", true}, false, []string{"cmd", "--pretty", "--separator=semicolon", "test.csv"}},
		{"Separator not identified", model.InputFile{}, true, []string{"cmd", "--separator=pipe", "test.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOsArgs := os.Args
			defer func() {
				os.Args = actualOsArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()

			os.Args = tt.osArgs

			got, err := GetFileData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckIfValidFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test*.csv")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	tests := []struct {
		name     string
		filename string
		want     bool
		wantErr  bool
	}{
		{"File does exist", tmpfile.Name(), true, false},
		{"File does not exist", "nowhere/test.csv", false, true},
		{"File is not csv", "test.txt", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckIfValidFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIfValidFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckIfValidFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
