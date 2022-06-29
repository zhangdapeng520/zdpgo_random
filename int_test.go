package zdpgo_random

import (
	"testing"
)

// 生成随机的整数
func TestRandomInt(t *testing.T) {
	r := New()
	var data []int
	for i := 0; i < 10; i++ {
		data = append(data, r.Int(0, 100))
		if i > 1 {
			if data[i] == data[i-1] {
				panic("随机数重复")
			}
		}
	}
}

func TestRandom_Int32(t *testing.T) {
	r := New()
	var data []int32
	for i := 0; i < 10; i++ {
		data = append(data, r.Int32(10000000, 99999999))
		if i > 1 {
			if data[i] == data[i-1] {
				panic("随机数重复")
			}
		}
	}
}

func TestRandom_Int64(t *testing.T) {
	r := New()
	var data []int64
	for i := 0; i < 10; i++ {
		data = append(data, r.Int64(10000000, 99999999))
		if i > 1 {
			if data[i] == data[i-1] {
				panic("随机数重复")
			}
		}
	}
}
