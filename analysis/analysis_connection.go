package analysis

import (
	"github.com/dalebao/gRedis-cli/pkg"
	"github.com/dalebao/gRedis-cli/util"
	"strconv"
	"strings"
)

func Analysis(key string) string {
	res, _ := pkg.Client.DebugObject(key).Result()

	r := strings.Split(res, " ")
	s := strings.Split(r[4], ":")

	a, _ := strconv.ParseUint(s[1], 10, 64)
	return util.HumanSize(a)
}
