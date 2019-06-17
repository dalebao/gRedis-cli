package pkg

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/bndr/gotabulate"
	"github.com/dalebao/gRedis-cli/export"
	"strconv"
	"time"
)

/**
redis keys 命令
 */
func HandleCmdKey(params []string, e map[string]string) {
	header := []string{"Type", "Key", "TTL"}
	res, _ := ScanKeys(0, params[0])
	keys := &Keys{}
	keys.Set(e)
	data := keys.DiffType(res)
	if keys.Export != "" {
		uExport := &export.UExport{FileName: "keys-" + params[0], Data: data, Type: keys.Export, Header: header}
		fileName, err := uExport.Export()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("导出结果成功，文件名为：" + fileName)
	} else {
		printTable(data, header)
	}
}

/**
redis 获取键内容 命令
目前支持
string=>get
hash=>hgetall
 */
func HandleCmdGet(params []string, e map[string]string) {
	rKey := params[0]
	rType, _ := Client.Type(rKey).Result()
	fmt.Println(rType)
	switch rType {
	case "none":
		fmt.Println(rKey + "不存在")
	case "string":
		r, _ := Client.Get(rKey).Result()
		fmt.Println(r)
	case "hash":
		var data [][]string
		r, _ := Client.HGetAll(rKey).Result()
		for k, v := range r {
			info := []string{k, v}
			data = append(data, info)
		}
		printTable(data, []string{"Key", "Value"})
	case "list":
		r, _ := Client.LRange(rKey, 0, -1).Result()
		var data [][]string
		for _, v := range r {
			info := []string{v}
			data = append(data, info)
		}
		printTable(data, []string{"Value", "left->right"})
	case "set":
		r, _ := Client.SMembers(rKey).Result()
		var data [][]string
		for _, v := range r {
			info := []string{v}
			data = append(data, info)
		}
		printTable(data, []string{"Value"})

	case "zset":
		r, _ := Client.ZRangeWithScores(rKey, 0, -1).Result()
		var data [][]string
		for _, v := range r {
			member := fmt.Sprintf("%v", v.Member)
			score := fmt.Sprintf("%f", v.Score)
			info := []string{member, score}
			data = append(data, info)
		}
		printTable(data, []string{"Member", "Score"})
	}

}

/**
查询多个redis键的类型
 */
func HandleCmdType(params []string, e map[string]string) {
	kLen := len(params)
	var data [][]string
	for i := 0; i < kLen; i++ {
		rKey := params[i]
		rType, _ := Client.Type(rKey).Result()

		info := []string{rKey, rType}
		data = append(data, info)
	}

	printTable(data, []string{"Key", "Type"})
}

/**
查询多个redis键的ttl
 */
func HandleCmdTTL(params []string, e map[string]string) {
	kLen := len(params)
	var data [][]string
	for i := 0; i < kLen; i++ {
		rKey := params[i]
		rType, _ := Client.TTL(rKey).Result()

		info := []string{rKey, rType.String()}
		data = append(data, info)
	}

	printTable(data, []string{"Key", "TTL"})
}

/**
设置redis键过期时间
 */
func HandleCmdExpire(params []string, e map[string]string) {
	rKey := params[0]
	rExpire, _ := strconv.Atoi(params[1])
	Client.Expire(rKey, time.Duration(rExpire)*time.Second)

	var data [][]string
	rType, _ := Client.TTL(rKey).Result()
	info := []string{rKey, rType.String()}
	data = append(data, info)
	printTable(data, []string{"Key", "TTL"})
}

/**
删除redis键，多个删除
 */
func HandleCmdDel(params []string, e map[string]string) {
	rLen := len(params)

	var data [][]string
	for i := 0; i < rLen; i++ {
		rKey := params[i]
		res, _ := Client.Del(rKey).Result()
		r := "success"
		if res == 0 {
			r = "fail"
		}
		info := []string{rKey, r}
		data = append(data, info)
	}

	printTable(data, []string{"Key", "result"})
}

/**
使用通配符匹配redis键进行删除
 */
func HandleCmdRDel(params []string, e map[string]string) {
	re := params[0]
	res, _ := ScanKeys(0, re)
	rLen := len(res)
	show := false
	prompt := &survey.Confirm{
		Message: "共匹配到" + strconv.Itoa(rLen) + "条数据，Y选择删除，N直接删除",
	}
	survey.AskOne(prompt, &show)

	if show {
		sK := []string{}
		prompt1 := &survey.MultiSelect{
			Message: "请选择你想删除的键",
			Options: res,
		}
		survey.AskOne(prompt1, &sK)
		HandleCmdDel(sK, e)
		return
	}
	HandleCmdDel(res, e)
	return
}

/**
打印表格
 */
func printTable(data [][]string, k []string) {
	t := gotabulate.Create(data)

	t.SetHeaders(k)
	t.SetAlign("right")
	fmt.Println(t.Render("grid"))
}

func ScanKeys(cursor uint64, re string) ([]string, uint64) {

	res, c, _ := Client.Scan(cursor, re, 100).Result()
	if c != 0 {
		r, _ := ScanKeys(c, re)
		res = append(res, r...)
	}
	return res, c
}
