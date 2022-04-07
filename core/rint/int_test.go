package rint

import (
	"fmt"
	"testing"
)

func getInt() *Int {
	return NewInt()
}

// 生成随机的整数
func TestRandomInt(t *testing.T) {
	i1 := getInt()
	for i := 0; i < 10; i++ {
		fmt.Println(i1.Int(0, 100))
	}
}
func TestRandom_Int32(t *testing.T) {
	i1 := getInt()
	for i := 0; i < 10; i++ {
		fmt.Println(i1.Int32(10000000, 99999999))
	}
}

func TestRandom_Int64(t *testing.T) {
	i1 := getInt()
	for i := 0; i < 10; i++ {
		fmt.Println(i1.Int64(10000000, 99999999))
	}
}
