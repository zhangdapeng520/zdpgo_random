package zdpgo_random

import (
	"net"
)

// HttpPort 获取系统中可用的端口号
func (r *Random) HttpPort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		r.Log.Error("解析TCP地址失败", "error", err)
		return 0
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		r.Log.Error("创建tcp监听失败", "error", err)
		return 0
	}
	defer l.Close()

	// 获取端口号
	p := l.Addr().(*net.TCPAddr).Port
	r.Log.Debug("获取可用的HTTP端口号成功", "port", p)
	return p
}
