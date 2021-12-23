# zgo_random

生成随机数据：https://github.com/zhangdapeng520/zdpgo_random

功能列表
- 生成随机的，可用的HTTP端口号
- 生成随机的UUID字符串
- 生成随机的整数
- 生成随机的整数切片

官方教程
- zdpgo_docker框架使用教程(官方博客)：https://blog.csdn.net/qq_37703224/category_11546451.html
- zdpgo_docker框架使用教程(知乎视频)：https://www.zhihu.com/people/zhang-da-peng-24-84/zvideos
- zdpgo_docker框架使用教程(头条视频)：https://www.toutiao.com/c/user/token/MS4wLjABAAAAYjNW4m-DnKcdysZrnud9Jk8G4Z44gNUTK1u7tvz_XWg/??tab=video
- zdpgo_docker框架使用教程(BiliBili视频)：https://space.bilibili.com/384418921

## 一、快速入门

### 1.1 生成随机可用的端口号
```go
package main

import (
	`fmt`
	
	`github.com/zhangdapeng520/zdpgo_random`
)

func main() {
	r := zdpgo_random.New()
	
	// 生成随机可用的端口号
	for i := 0; i < 20; i++ {
		fmt.Println(r.RandomHttpPort())
	}
}
```

### 1.2 生成随机的UUID字符串
```go
package main

import (
	`fmt`
	`reflect`
	
	`github.com/zhangdapeng520/zdpgo_random`
)

func main() {
	r := zdpgo_random.New()
	for i := 0; i < 10; i++ {
		fmt.Println(r.RandomUUID(), reflect.TypeOf(r.RandomUUID()))
	}
}
```

### 1.3 生成随机的整数
```go
package main

import (
	`fmt`
	
	`github.com/zhangdapeng520/zdpgo_random`
)

func main() {
	r := zdpgo_random.New()
	fmt.Println(r.RandomInt(30, 40))
	fmt.Println(r.RandomInt(10, 20))
	fmt.Println(r.RandomInt(-10, 0))
}
```

### 1.4 生成随机的整数切片
```go
package main

import (
	`fmt`
	
	`github.com/zhangdapeng520/zdpgo_random`
)

func main() {
	r := zdpgo_random.New()
	fmt.Println(r.RandomIntSlice(0, 100, 1))
	fmt.Println(r.RandomIntSlice(0, 10, 10))
	fmt.Println(r.RandomIntSlice(-10, 10, 20))
}
```