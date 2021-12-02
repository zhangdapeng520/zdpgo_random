package http

import (
	"fmt"
	"testing"
)

func TestFreePort(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Println(FreePort())
	}
}
