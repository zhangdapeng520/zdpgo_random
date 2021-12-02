package uuid

import (
	"fmt"
	"testing"
)

// 测试解析uuid
func TestParseUUID(t *testing.T) {
	u1 := NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	// 从输入中解析UUID
	u2, err := FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("解析UUID失败: %s", err)
	}
	fmt.Printf("成功解析UUID: %s", u2)
}

// 测试生成UUID
func TestGeneratorUUID(t *testing.T) {
	fmt.Println(NewV1())
	fmt.Println(NewV4())
}
