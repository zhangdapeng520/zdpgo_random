package zdpgo_random

import (
	"github.com/zhangdapeng520/zdpgo_random/core/rint"
	"github.com/zhangdapeng520/zdpgo_random/core/rslice"
	"github.com/zhangdapeng520/zdpgo_random/core/str"
	"github.com/zhangdapeng520/zdpgo_random/core/user"
)

// Random 随机数据生成器核心对象
type Random struct {
	config *RandomConfig // 配置对象
	Int    *rint.Int     // 生成随机的整数
	Slice  *rslice.Slice // 生成随机的切片
	Str    *str.Str      // 生成随机的字符串
	User   *user.User    // 生成随机的用户数据
}

// New 生成随机数据对象
func New() *Random {
	return NewWithConfig(RandomConfig{})
}

// NewWithConfig 根据配置生成随机数对象
func NewWithConfig(config RandomConfig) *Random {
	r := Random{}

	// 指定配置
	r.config = &config

	// 随机对象
	r.Int = rint.NewInt()
	r.Slice = rslice.NewSlice()
	r.Str = str.NewStr()
	r.User = user.NewUser()

	// 返回随机对象
	return &r
}
