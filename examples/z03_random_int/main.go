package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_random"
)

func main() {
	r := zdpgo_random.New()
	fmt.Println(r.RandomInt(30, 40))
	fmt.Println(r.RandomInt(10, 20))
	fmt.Println(r.RandomInt(-10, 0))
}
