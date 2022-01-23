package zdpgo_random

// RandomIntSlice 生成随机的整数切片
// @param min：最小值
// @param max：最大值
// @param length：切片长度
func (r *Random) RandomIntSlice(min, max, length int) []int {
	var result []int
	for i := 0; i < length; i++ {
		random_int := r.Int(min, max)
		result = append(result, random_int)
	}
	return result
}
