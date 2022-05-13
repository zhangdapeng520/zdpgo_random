package zdpgo_random

import (
	"fmt"
	"testing"
)

// 生成随机的手机号
func TestRandom_Phone(t *testing.T) {
	u := getRandom()
	for i := 0; i < 10; i++ {
		fmt.Println(u.Phone())
	}
}

// 生成随机的中文名
func TestRandom_Name(t *testing.T) {
	u := getRandom()
	for i := 0; i < 100; i++ {
		fmt.Println(u.Name(true))
	}
	for i := 0; i < 100; i++ {
		fmt.Println(u.Name(false))
	}
}

// 生成随机的英文名
func TestRandom_EnglishName(t *testing.T) {
	u := getRandom()
	for i := 0; i < 100; i++ {
		fmt.Println(u.EnglishName())
	}
}

// 生成随机的邮箱
func TestRandom_Email(t *testing.T) {
	u := getRandom()
	for i := 0; i < 100; i++ {
		fmt.Println(u.Email())
	}
}
