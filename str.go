package zdpgo_random

import (
	"math/rand"
	"strconv"
	"strings"
)

// Str 生成指定长度的随机字符串
// 掩码加强版：rand.Int63会产生63bit的随机数，如果我们把它分成6份，那么一次就可以产生10个6bit的随机数。这样就减少了浪费。
func (s *Random) Str(n int) string {
	// 数据源
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	var src = rand.NewSource(rand.Int63())
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// IntStr 生成随机长度的数字
func (s *Random) IntStr(n int) string {
	// 检查参数合法性
	if n <= 0 {
		return "0"
	}

	// 生成随机数字组成的字符串
	var arr []string
	for i := 0; i < n; i++ {
		num := rand.Intn(10)
		numS := strconv.Itoa(num)
		arr = append(arr, numS)
	}
	result := strings.Join(arr, "")
	return result
}
