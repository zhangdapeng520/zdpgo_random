package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_random"
)

func main() {
	r := zdpgo_random.New()

	// 生成随机可用的端口号
	for i := 0; i < 20; i++ {
		fmt.Println(r.RandomHttpPort())
	}
}
