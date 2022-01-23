package zdpgo_random

import "github.com/zhangdapeng520/zdpgo_log"

// Random 随机数据生成器核心对象
type Random struct {
	log    *zdpgo_log.Log // 日志对象
	config *RandomConfig  // 配置对象
}

// RandomConfig 随机生成器配置对象
type RandomConfig struct {
	Debug       bool   // 是否为debug模式
	LogFilePath string // 日志路径
}

// SetDebug 设置debug模式
func (r *Random) SetDebug(debug bool) {
	r.config.Debug = debug
}

// IsDebug 是否为debug模式
func (r *Random) IsDebug() bool {
	return r.config.Debug
}

// SetLogPath 设置日志路径
func (r *Random) SetLogPath(logPath string) {
	r.config.LogFilePath = logPath
	logConfig := zdpgo_log.LogConfig{
		Debug: r.config.Debug,
		Path:  r.config.LogFilePath,
	}
	l := zdpgo_log.New(logConfig)
	r.log = l
}

// New 生成随机数据对象
func New(config RandomConfig) *Random {
	r := Random{}

	// 生成日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_random.log"
	}
	logConfig := zdpgo_log.LogConfig{
		Debug: config.Debug,
		Path:  config.LogFilePath,
	}
	l := zdpgo_log.New(logConfig)
	r.log = l

	// 指定配置
	r.config = &config
	return &r
}
