package zdpgo_random

import "github.com/zhangdapeng520/zdpgo_log"

/*
@Time : 2022/5/13 10:32
@Author : 张大鹏
@File : random_test.go
@Software: Goland2021.3.1
@Description:
*/

func getRandom() *Random {
	return New(zdpgo_log.Tmp)
}
