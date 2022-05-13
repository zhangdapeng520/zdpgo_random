package zdpgo_random

import (
	"fmt"
	"testing"
)

// 测试生成随机的切片
func TestRandomIntSlice(t *testing.T) {
	s := getRandom()

	// 生成10个随机数组
	for i := 0; i < 10; i++ {
		fmt.Println(s.SliceInt(0, 10, 10))
	}
}
