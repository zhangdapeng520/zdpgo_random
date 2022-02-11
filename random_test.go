package zdpgo_random

import (
	"fmt"
	"testing"
)

func prepareRandom() *Random {
	r := New(RandomConfig{
		Debug: true,
	})
	return r
}

func TestRandom_New(t *testing.T) {
	r := prepareRandom()
	fmt.Println(r)
}
