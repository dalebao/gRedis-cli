package analysis

import (
	"fmt"
	"github.com/dalebao/gRedis-cli/pkg"
	"testing"
)

func init() {
	r := pkg.RedisConfig{Addr: "192.168.20.247", Port: "6379", Password: "51cartest1234"}
	r.Dial()
}

func TestAnalysis(t *testing.T) {
	s := Analysis("entry_inventory_90_to_180_day_v_24")
	fmt.Println(s)
}
