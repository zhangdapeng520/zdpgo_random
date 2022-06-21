package zdpgo_random

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"testing"
)

/*
@Time : 2022/6/21 13:53
@Author : 张大鹏
@File : article_test
@Software: Goland2021.3.1
@Description:
*/

func TestRandom_RandomArticle(t *testing.T) {
	r := New(zdpgo_log.Tmp)
	//article := r.RandomArticle()
	//fmt.Println(article.Title, article.Author, article.Content)
	for i := 0; i < 1000; i++ {
		article := r.RandomArticle()
		if article.Title == "" {
			panic("文章标题为空")
		}
		if article.Author == "" {
			panic("文章作者为空")
		}
		if article.Content == "" {
			panic("文章内容为空")
		}
	}
}
