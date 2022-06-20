package zdpgo_random

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"testing"
)

// 测试生成随机的切片
func TestRandomIntSlice(t *testing.T) {
	r := New(zdpgo_log.Tmp)

	// 生成10个随机数组
	for i := 0; i < 10; i++ {
		fmt.Println(r.SliceInt(0, 10, 10))
	}
}
