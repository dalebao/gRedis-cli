package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"sync"
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

	data := make(chan []string)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	for _, r := range exportCsv.Data {
		wg.Add(1)
		go func() {
			data <- r
			wg.Done()
		}()
	}

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write(exportCsv.Header)

	go func() {
		for line := range data {
			w.Write(line)
		}
		done <- true
	}()

	go func() {
		wg.Wait()
		close(data)
	}()

	if <-done {
		w.Flush()
	}
	w.Flush()
	return fileName, nil
}

func generateAddr(fileName string) (addr string) {
	t := time.Now()
	addr = fileName + "-" + t.Format(time.ANSIC) + ".csv"
	return
}
