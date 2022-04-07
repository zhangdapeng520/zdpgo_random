package rslice

import "github.com/zhangdapeng520/zdpgo_random/core/rint"

type Slice struct {
	int *rint.Int
}

func NewSlice() *Slice {
	s := Slice{}
	s.int = rint.NewInt()
	return &s
}

// IntSlice 生成随机的整数切片
// @param min：最小值
// @param max：最大值
// @param length：切片长度
func (s *Slice) IntSlice(min, max, length int) []int {
	var result []int
	for i := 0; i < length; i++ {
		randomInt := s.int.Int(min, max)
		result = append(result, randomInt)
	}
	return result
}
