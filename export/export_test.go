package export

import (
	"fmt"
	"testing"
)

func TestExportGenerator(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	v := &UExport{FileName: "keys_*_export.csv", Data: records, Header: []string{"helo", "hi", "test"}, Type: "csv"}
	fileName, err := v.Export()
	fmt.Println(fileName, err)
}
