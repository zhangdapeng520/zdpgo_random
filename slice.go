package zdpgo_random

// 生成随机的整数切片
func RandomIntSlice(min, max, length int) []int {
	var result []int
	for i := 0; i < length; i++ {
		random_int := RandomInt(min, max)
		result = append(result, random_int)
	}
	return result
}
