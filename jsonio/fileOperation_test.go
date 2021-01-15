package jsonio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWriteJSONFile(t *testing.T) {
	dataMap := []map[string]string{
		{"COL1": "1", "COL2": "2", "COL3": "3"},
		{"COL1": "4", "COL2": "5", "COL3": "6"},
	}
	tests := []struct {
		csvPath  string
		jsonPath string
		pretty   bool
		name     string
	}{
		{"notPrettyJson.csv", "notPrettyJson.json", false, "notPretty JSON"},
		{"prettyJson.csv", "prettyJson.json", true, "Pretty JSON"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writerChannel := make(chan map[string]string)
			done := make(chan bool)
			go func() {
				for _, record := range dataMap {
					writerChannel <- record
				}
				close(writerChannel)
			}()
			go WriteJSONFile(tt.csvPath, writerChannel, done, tt.pretty)
			<-done
			testOutput, err := ioutil.ReadFile(tt.jsonPath)

			if err != nil { // Failing test if something went wrong with our JSON file creation
				t.Errorf("writeJSONFile(), Output file got error: %v", err)
			}
			defer os.Remove(tt.jsonPath)
			wantOutput, err := ioutil.ReadFile(filepath.Join("testing", tt.jsonPath))
			if err != nil {
				panic(err)
			}
			if !reflect.DeepEqual(testOutput, testOutput) {
				t.Errorf("writeJSONFile() = \n%v, want \n%v", string(testOutput), string(wantOutput))
			}
		})
	}
}
