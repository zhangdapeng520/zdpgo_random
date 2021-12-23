package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_random"
)

func main() {
	r := zdpgo_random.New()
	fmt.Println(r.RandomIntSlice(0, 100, 1))
	fmt.Println(r.RandomIntSlice(0, 10, 10))
	fmt.Println(r.RandomIntSlice(-10, 10, 20))
}
