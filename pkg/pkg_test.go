package pkg

import (
	"github.com/dalebao/gRedis-cli/handler"
	"testing"
)

func init() {
	existsConfig := RedisConfig{Addr: "127.0.0.1", Port: "6379"}
	existsConfig.Dial()

}

func TestHandleCmdKey(t *testing.T) {
	params := []string{"*"}
	e := make(map[string]string)
	handler.HandleCmdKey(params, e)
}
