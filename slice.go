package zdpgo_random

// SliceInt 生成随机的整数切片
// @param min：最小值
// @param max：最大值
// @param length：切片长度
func (s *Random) SliceInt(min, max, length int) []int {
	var result []int
	for i := 0; i < length; i++ {
		randomInt := s.Int(min, max)
		result = append(result, randomInt)
	}
	return result
}
