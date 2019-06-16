package export

import (
	"errors"
	"github.com/dalebao/gRedis-cli/export/csv"
)

type Export interface {
	Generator() (string, error)
}

type UExport struct {
	FileName string
	Type     string
	Data     [][]string
	Header   []string
}

func (uExport *UExport) Export() (string, error) {
	var export Export
	var err error
	switch uExport.Type {
	case "csv":
		export = &csv.ExportCsv{FileName: uExport.FileName, Data: uExport.Data, Header: uExport.Header}
	default:
		err = errors.New(uExport.Type + "格式不支持")
	}
	fileName, err := export.Generator()
	if err != nil {
		return "", err
	}
	return fileName, nil
}
