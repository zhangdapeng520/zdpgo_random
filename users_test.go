package zdpgo_random

import (
	"fmt"
	"testing"
)

func TestRandom_Phone(t *testing.T) {
	r := prepareRandom()
	for i := 0; i < 10; i++ {
		fmt.Println(r.Phone())
	}
}

func TestRandom_Name(t *testing.T) {
	r := prepareRandom()
	for i := 0; i < 100; i++ {
		fmt.Println(r.Name(true))
	}
	for i := 0; i < 100; i++ {
		fmt.Println(r.Name(false))
	}
}

func TestRandom_EnglishName(t *testing.T) {
	r := prepareRandom()
	for i := 0; i < 100; i++ {
		fmt.Println(r.EnglishName())
	}
}

func TestRandom_Email(t *testing.T) {
	r := prepareRandom()
	for i := 0; i < 100; i++ {
		fmt.Println(r.Email())
	}
}
