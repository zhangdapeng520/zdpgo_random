package zdpgo_random

import (
	"fmt"
	"testing"
)

// 生成随机的整数切片
func TestRandomIntSlice(t *testing.T) {
	arr := RandomIntSlice(0, 100, 20)
	fmt.Println(arr)
}
