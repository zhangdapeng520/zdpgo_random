package zdpgo_random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // 随机的种子
}

// Random 随机数据生成器核心对象
type Random struct {
	Config *Config // 配置对象
}

// New 生成随机数据对象
func New() *Random {
	return NewWithConfig(&Config{})
}

// NewWithConfig 通过配置生成随机数据对象
func NewWithConfig(config *Config) *Random {
	r := Random{}

	// 指定配置
	r.Config = config

	// 返回随机对象
	return &r
}
