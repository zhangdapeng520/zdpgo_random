package zdpgo_random

import (
	"testing"
)

// 测试随机字符串的生成
func TestStr_str(t *testing.T) {
	r := New()
	var data []string
	for i := 0; i < 1000; i++ {
		tmp := r.Str(16)
		data = append(data, tmp)
		if i > 1 {
			if data[i] == data[i-1] {
				panic("随机字符串重复")
			}
		}
	}
}

// 测试生成指定长度的随机数字符串
func TestStr_intStr(t *testing.T) {
	r := New()
	var data []string
	for i := 0; i < 1000; i++ {
		tmp := r.IntStr(16)
		data = append(data, tmp)
		if i > 1 {
			if data[i] == data[i-1] {
				panic("随机字符串重复")
			}
		}
	}
}
