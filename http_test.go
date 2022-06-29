package zdpgo_random

import (
	"testing"
)

// 获取系统中可用的端口号
func TestRandomHttpPort(t *testing.T) {
	r := New()

	var ports []int
	for i := 0; i < 10; i++ {
		ports = append(ports, r.HttpPort())
		if i > 1 {
			if ports[i] == ports[0] {
				panic("端口号重复")
			}
		}
	}
}
