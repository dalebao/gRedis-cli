package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/dalebao/gRedis-cli/pkg"
	"strings"
)

var simpleQs = []*survey.Question{
	{
		Name: "addr",
		Prompt: &survey.Input{
			Message: "redis-server address?",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "Port",
		Prompt: &survey.Input{
			Message: "redis-server port?",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "Password",
		Prompt: &survey.Input{
			Message: "redis-server password?",
		},
		Transform: survey.Title,
	},
	{
		Name: "db",
		Prompt: &survey.Input{
			Message: "redis-server db?",
		},
		Transform: survey.Title,
	},
}

func main() {
	err := choseConfig()
	for err != nil {
		tryAgain := true
		prompt := &survey.Confirm{
			Message: "是否重新选择",
		}
		survey.AskOne(prompt, &tryAgain)
		if tryAgain {
			err = choseConfig()
		} else {
			err = nil
		}
	}

	for {
		name := ""
		prompt := &survey.Input{
			Message: "请输入命令:",
		}
		survey.AskOne(prompt, &name, survey.WithValidator(survey.Required))
		r, e := HandleCmd(strings.TrimSpace(name))
		invokeCmd(r, e)

		if r[0] == "quit" {
			fmt.Println("Bye~ Bye!!")
			break
		}

	}

}

func handleCmd(name string) []string {
	return strings.Split(name, " ")
}

func HandleCmd(name string) ([]string, map[string]string) {
	all := strings.Split(name, "-")
	if len(all) > 2 {
		fmt.Println("额外参数隔断 - 只需一个")
	}
	r := strings.Split(all[0], " ")
	eCmd := make(map[string]string)

	if len(all) == 2 {
		for _, v := range strings.Split(all[1], " ") {
			n := strings.Split(v, "=")
			eCmd[n[0]] = n[1]
		}
	}

	return r, eCmd
}

/**
解析命令
 */
func invokeCmd(r []string, e map[string]string) {
	cmd := r[0]
	p := r[1:]
	switch cmd {
	case "keys":
		pkg.HandleCmdKey(p, e)
	case "get":
		pkg.HandleCmdGet(p, e)
	case "type":
		pkg.HandleCmdType(p, e)
	case "ttl":
		pkg.HandleCmdTTL(p, e)
	case "expire":
		pkg.HandleCmdExpire(p, e)
	case "del":
		pkg.HandleCmdDel(p, e)
	case "rdel":
		pkg.HandleCmdRDel(p, e)
	default:
		fmt.Println("命令不存在")
	}
}

/**
选择配置文件
 */
func choseConfig() error {
	zc := pkg.ZC{}
	rConfig := pkg.RedisConfig{}

	var choseOne string

	existsConfig := zc.ReadConfig()
	prompt := &survey.Select{
		Message: "请选择你的配置",
		Options: existsConfig,
	}
	survey.AskOne(prompt, &choseOne)

	if choseOne == "手动输入" {
		if rConfig == (pkg.RedisConfig{}) {
			err := survey.Ask(simpleQs, &rConfig)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = rConfig.Dial()
			if err != nil {
				fmt.Println("链接失败")
				return err
			}
			successSave := false
			for !successSave {
				isSave := false
				prompt := &survey.Confirm{
					Message: "链接成功是否保存",
				}
				survey.AskOne(prompt, &isSave)
				if isSave {
					configName := ""
					prompt := &survey.Input{
						Message: "请输入你的配置名称",
					}
					survey.AskOne(prompt, &configName, survey.WithValidator(survey.Required))
					err = zc.SaveConfig(configName, rConfig)
					if err != nil {
						fmt.Println(err)
					} else {
						successSave = true
					}
				} else {
					successSave = true
				}
			}
		}
	} else {
		rConfig = zc.FindConfig(choseOne)
		err := rConfig.Dial()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("链接成功")
	}
	return nil
}
