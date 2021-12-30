package zdpgo_random

import (
	"fmt"
	"testing"
)

// 生成随机的UUID字符串（RFC 4122）
func TestRandomUUID(t *testing.T) {
	r := New()
	for i := 0; i < 100; i++ {
		fmt.Println(r.RandomUUID())
	}
}
