package zdpgo_random

import (
	"testing"
)

/*
@Time : 2022/6/21 13:57
@Author : 张大鹏
@File : random_test
@Software: Goland2021.3.1
@Description:
*/

func TestRandom_New(t *testing.T) {
	r := New()
	if r == nil {
		panic(r)
	}

	r = NewWithConfig(&Config{})
	if r == nil {
		panic(r)
	}
}
