package main

import (
	"fmt"
	"reflect"

	"github.com/zhangdapeng520/zdpgo_random"
)

func main() {
	r := zdpgo_random.New()
	for i := 0; i < 10; i++ {
		fmt.Println(r.RandomUUID(), reflect.TypeOf(r.RandomUUID()))
	}
}
