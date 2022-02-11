package zdpgo_random

import (
	"github.com/zhangdapeng520/zdpgo_zap"
)

// Random 随机数据生成器核心对象
type Random struct {
	log    *zdpgo_zap.Zap // 日志对象
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

// New 生成随机数据对象
func New(config RandomConfig) *Random {
	r := Random{}

	// 生成日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_random.log"
	}
	r.log = zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        config.Debug,
		OpenGlobal:   true,
		OpenFileName: true,
		LogFilePath:  config.LogFilePath,
	})

	// 指定配置
	r.config = &config
	return &r
}
