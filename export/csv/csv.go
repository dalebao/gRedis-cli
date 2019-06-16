package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"time"
)

type ExportCsv struct {
	FileName string
	Data     [][]string
	Header   []string
}

func (exportCsv *ExportCsv) Generator() (string, error) {
	fileName := generateAddr(exportCsv.FileName)
	f, err := os.Create(fileName)
	if (err != nil) {
		return "", errors.New("创建文件失败")
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write(exportCsv.Header)
	for _, line := range exportCsv.Data {
		w.Write(line)
	}
	w.Flush()

	return fileName, nil
}

func generateAddr(fileName string) (addr string) {
	t := time.Now()
	addr = fileName + "-" + t.Format(time.ANSIC) + ".csv"
	return
}
