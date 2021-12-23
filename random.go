package zdpgo_random

import "github.com/zhangdapeng520/zdpgo_log"

type Random struct {
	logPath string         // 日志路径
	log     *zdpgo_log.Log // 日志对象
	debug   bool           // 是否为debug模式
}

// 设置debug模式
func (r *Random) SetDebug(debug bool) {
	r.debug = debug
}

// 设置日志路径
func (r *Random) SetLogPath(logPath string) {
	r.logPath = logPath
	r.log = zdpgo_log.New(logPath)
	r.log.SetDebug(r.debug)
}

// 生成随机对象
func New() *Random {
	r := Random{}
	r.SetDebug(false)                // 默认为非debug模式
	r.SetLogPath("zdpgo_random.log") // 默认的日志路径
	return &r
}
