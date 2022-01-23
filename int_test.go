package zdpgo_random

import (
	"fmt"
	"testing"
)

// 生成随机对象
func prepareRandom() *Random {
	r := New(RandomConfig{
		Debug: true,
	})
	return r
}

// 生成随机的整数
func TestRandomInt(t *testing.T) {
	r := prepareRandom()
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int(0, 100))
	}
}
func TestRandom_Int32(t *testing.T) {
	r := prepareRandom()
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int32(10000000, 99999999))
	}
	
	var num int32= r.Int32(10000000, 99999999)
	fmt.Println(num)
}

func TestRandom_Int64(t *testing.T) {
	r := prepareRandom()
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int64(10000000, 99999999))
	}
}
