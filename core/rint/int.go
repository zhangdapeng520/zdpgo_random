package rint

import (
	"math/rand"
	"time"
)

type Int struct {
}

func NewInt() *Int {
	i := Int{}
	rand.Seed(time.Now().UnixNano())
	return &i
}

// Int 生成随机的整数
func (i *Int) Int(min, max int) int {
	return rand.Intn(max-min) + min
}

// Int32 生成随机的int32整数
func (i *Int) Int32(min, max int32) int32 {
	start := max - min
	result := rand.Int31n(start) + min
	return result
}

// Int64 生成随机的int64整数
func (i *Int) Int64(min, max int64) int64 {
	start := max - min
	result := rand.Int63n(start) + min
	return result
}
