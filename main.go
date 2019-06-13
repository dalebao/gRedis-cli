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
	rConfig := pkg.RedisConfig{}
	if rConfig == (pkg.RedisConfig{}) {
		err := survey.Ask(simpleQs, &rConfig)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = rConfig.Dial()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("链接成功")
	}

	for {
		name := ""
		prompt := &survey.Input{
			Message: "请输入命令:",
		}
		survey.AskOne(prompt, &name, survey.WithValidator(survey.Required))
		r := handleCmd(strings.TrimSpace(name))

		invokeCmd(r)

		if r[0] == "quit" {
			fmt.Println("Bye~ Bye!!")
			break
		}

	}

}

func handleCmd(name string) []string {
	return strings.Split(name, " ")
}

func invokeCmd(r []string) {
	switch r[0] {
	case "keys":
		pkg.HandleCmdKey(r)
	case "get":
		pkg.HandleCmdGet(r)
	case "type":
		pkg.HandleCmdType(r)
	case "ttl":
		pkg.HandleCmdTTL(r)
	case "expire":
		pkg.HandleCmdExpire(r)
	case "del":
		pkg.HandleCmdDel(r)
	case "rdel":
		pkg.HandleCmdRDel(r)


	}
}
