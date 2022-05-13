package zdpgo_random

import (
	"math/rand"
)

// Int 生成随机的整数
func (i *Random) Int(min, max int) int {
	return rand.Intn(max-min) + min
}

// Int32 生成随机的int32整数
func (i *Random) Int32(min, max int32) int32 {
	start := max - min
	result := rand.Int31n(start) + min
	return result
}

// Int64 生成随机的int64整数
func (i *Random) Int64(min, max int64) int64 {
	start := max - min
	result := rand.Int63n(start) + min
	return result
}
