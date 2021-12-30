package zdpgo_random

import (
	"fmt"
	"testing"
)

// 生成随机的整数
func TestRandomInt(t *testing.T) {
	r := New()
	for i := 0; i < 10; i++ {
		fmt.Println(r.RandomInt(0, 100))
	}
}
