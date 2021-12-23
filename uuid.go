package zdpgo_random

import (
	"crypto/rand"
	"fmt"
)

// 生成随机的UUID字符串（RFC 4122）
func (r *Random) RandomUUID() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		r.log.Error("生成UUID失败：", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid
}
