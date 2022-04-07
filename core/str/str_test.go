package str

import "testing"

func getStr() *Str {
	return NewStr()
}

// 测试随机字符串的生成
func TestStr_str(t *testing.T) {
	s := getStr()
	for i := 0; i < 1000; i++ {
		t.Log(s.Str(i))
	}
}
