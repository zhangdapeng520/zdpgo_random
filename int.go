package zdpgo_random

import (
	"math/rand"
)

// 生成随机的整数
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}