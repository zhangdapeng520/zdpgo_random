package zdpgo_random

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // 随机的种子
}

var (
	Log *zdpgo_log.Log
)

// Random 随机数据生成器核心对象
type Random struct {
	Config *Config // 配置对象
}

// New 生成随机数据对象
func New(log *zdpgo_log.Log) *Random {
	return NewWithConfig(&Config{}, log)
}

// NewWithConfig 通过配置生成随机数据对象
func NewWithConfig(config *Config, log *zdpgo_log.Log) *Random {
	r := Random{}

	// 日志
	Log = log

	// 指定配置
	r.Config = config

	// 返回随机对象
	return &r
}
