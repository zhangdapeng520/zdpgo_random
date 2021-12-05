package zdpgo_random

import (
	"fmt"
	"testing"
)

// 生成随机的整数
func TestRandomInt(t *testing.T) {
	fmt.Println(RandomInt(30, 40))
	fmt.Println(RandomInt(10, 20))
	fmt.Println(RandomInt(-10, 0))
}
