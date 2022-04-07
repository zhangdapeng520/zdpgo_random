package rslice

import (
	"fmt"
	"testing"
)

func getSlice() *Slice {
	s := Slice{}
	return &s
}

// 测试生成随机的切片
func TestRandomIntSlice(t *testing.T) {
	s := getSlice()
	// 生成10个随机数组
	for i := 0; i < 10; i++ {
		fmt.Println(s.IntSlice(0, 10, 10))
	}
}
