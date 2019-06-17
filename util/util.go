package util

import (
	"fmt"
	"strconv"
)

func HumanSize(byte uint64) string {
	units := []string{"Bytes", "KB", "MB", "GB", "TB", "PB", "EB"}
	fb := float64(byte)
	i := 0
	for ; fb >= 1024; i++ {
		fb /= 1024
	}
	size, _ := strconv.ParseFloat(fmt.Sprintf("%0.3f", fb), 64)
	return fmt.Sprintf("%0.3f %s", size, units[i])
}
