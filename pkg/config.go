package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type C struct {
	Name   string      `json:"name"`
	Config RedisConfig `json:"config"`
}

type ZC struct {
	Config []C `json:"config"`
}

const CONFIG_FILE  = ".config.json"

func init() {
	generateConfigFile()
}

/**
生成配置文件
 */
func generateConfigFile() {
	if !exists(CONFIG_FILE) {
		file, err := os.Create(CONFIG_FILE)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()
	}
	fmt.Println("file exists")
}

/**
判断文件是否存在
 */
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
将json数据转换为struct
 */
func (zc *ZC) ReadConfig() []string {
	zc.loadConfig()
	configName := []string{"手动输入"}

	for _, v := range zc.Config {
		configName = append(configName, v.Name)
	}

	return configName
}

/**
加载配置文件
 */
func (zc *ZC) loadConfig() {
	// 打开json文件
	fh, err := os.Open(CONFIG_FILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	// 读取json文件，保存到jsonData中
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(jsonData, &zc)
}

/**
查询配置
 */
func (zc *ZC) FindConfig(name string) RedisConfig {
	for _, v := range zc.Config {
		if v.Name == name {
			return v.Config
		}
	}
	return RedisConfig{}
}

/**
判断配置文件是否存在
 */
func (zc *ZC) existConfig(name string) bool {
	for _, v := range zc.Config {
		if v.Name == name {
			return true
		}
	}
	return false
}

/**
保存配置
 */
func (zc *ZC) SaveConfig(name string, config RedisConfig) error {
	if zc.existConfig(name) {
		return errors.New("配置" + name + "已存在")
	}
	c := C{Name: name, Config: config}
	zc.Config = append(zc.Config, c)
	zc.saveToFile()
	return nil
}

/**
将配置转换成json 保存到文件中
 */
func (zc *ZC) saveToFile() {
	configJson, err := json.Marshal(zc)
	if err != nil {
		fmt.Println(err)
	}
	jFile, err := os.OpenFile(CONFIG_FILE, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer jFile.Close()

	if _, err := jFile.Write(configJson); err != nil {
		fmt.Println(err)
	}
	fmt.Println("保存配置成功")
}
