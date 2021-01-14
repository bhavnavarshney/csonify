package csv

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/bhavnavarshney/csonify/model"
)

func Test_processCsvFile(t *testing.T) {
	wantMapSlice := []map[string]string{
		{"COL1": "1", "COL2": "2", "COL3": "3"},
		{"COL1": "4", "COL2": "5", "COL3": "6"},
	}
	tests := []struct {
		name      string
		csvString string
		separator string
	}{
		{"Comma separator", "COL1,COL2,COL3\n1,2,3\n4,5,6\n", "comma"},
		{"Semicolon separator", "COL1;COL2;COL3\n1;2;3\n4;5;6\n", "semicolon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpfile, err := ioutil.TempFile("", "test*.csv")
			if err != nil {
				panic(err)
			}
			defer os.Remove(tmpfile.Name())
			_, err = tmpfile.WriteString(tt.csvString)
			tmpfile.Sync()
			testFileData := model.InputFile{
				FilePath:  tmpfile.Name(),
				Pretty:    false,
				Separator: tt.separator,
			}
			writerChannel := make(chan map[string]string)
			go Process(testFileData, writerChannel)
			for _, wantMap := range wantMapSlice {
				record := <-writerChannel
				if !reflect.DeepEqual(record, wantMap) {
					t.Errorf("Process() = %v, want %v", record, wantMap)
				}
			}
		})
	}
}
