package zdpgo_random

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"testing"
)

// 生成随机的手机号
func TestRandom_Phone(t *testing.T) {
	r := New(zdpgo_log.Tmp)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Phone())
	}
}

// 生成随机的中文名
func TestRandom_Name(t *testing.T) {
	r := New(zdpgo_log.Tmp)
	for i := 0; i < 100; i++ {
		fmt.Println(r.Name(true))
	}
	for i := 0; i < 100; i++ {
		fmt.Println(r.Name(false))
	}
}

// 生成随机的英文名
func TestRandom_EnglishName(t *testing.T) {
	r := New(zdpgo_log.Tmp)
	for i := 0; i < 100; i++ {
		fmt.Println(r.EnglishName())
	}
}

// 生成随机的邮箱
func TestRandom_Email(t *testing.T) {
	r := New(zdpgo_log.Tmp)
	for i := 0; i < 100; i++ {
		fmt.Println(r.Email())
	}
}
