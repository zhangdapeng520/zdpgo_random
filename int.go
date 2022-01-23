package zdpgo_random

import (
	"math/rand"
	"time"
)

// RandomInt 生成随机的整数
func (r *Random) Int(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Int32 生成随机的int32整数
func (r *Random) Int32(min, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	start := max - min
	result := rand.Int31n(start) + min
	return result
}

// Int64 生成随机的int64整数
func (r *Random) Int64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	start := max - min
	result := rand.Int63n(start) + min
	return result
}
