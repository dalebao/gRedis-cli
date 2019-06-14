package test

import (
	"fmt"
	"github.com/dalebao/gRedis-cli/pkg"
	"testing"
)

func init() {
	rConfig := pkg.RedisConfig{"192.168.20.247", "6379", "51cartest1234", ""}
	err := rConfig.Dial()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestScanKeys(t *testing.T) {
	res, _ := pkg.ScanKeys(0, "v*")
	fmt.Println(res)
}

func TestConfigRead(t *testing.T) {
	zc := &pkg.ZC{}
	zc.ReadConfig()
}

func TestConfigSearch(t *testing.T) {
	zc := &pkg.ZC{}
	zc.ReadConfig()
	fmt.Println(zc.FindConfig("config1"))
}

func TestConfigConfigSave(t *testing.T) {
	zc := &pkg.ZC{}
	zc.ReadConfig()
	err := zc.SaveConfig("config3",pkg.RedisConfig{"192.168.20.247", "6379", "51cartest1234", ""})
	fmt.Println(err)
}
