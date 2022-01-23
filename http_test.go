package zdpgo_random

import (
	"fmt"
	"testing"
)

// 获取系统中可用的端口号
func TestRandomHttpPort(t *testing.T) {
	r := prepareRandom()
	for i := 0; i < 10; i++ {
		fmt.Println(r.RandomHttpPort())
	}
}
