package rhttp

import (
	"fmt"
	"testing"
)

func getHttp() *HTTP {
	return NewHTTP()

}

// 获取系统中可用的端口号
func TestRandomHttpPort(t *testing.T) {
	h := getHttp()
	for i := 0; i < 10; i++ {
		fmt.Println(h.Port())
	}
}
