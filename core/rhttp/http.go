package rhttp

import (
	"net"
)

type HTTP struct {
}

func NewHTTP() *HTTP {
	h := HTTP{}

	return &h
}

// Port 获取系统中可用的端口号
func (h *HTTP) Port() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}
