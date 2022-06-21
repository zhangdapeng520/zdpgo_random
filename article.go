package zdpgo_random

/*
@Time : 2022/6/21 13:32
@Author : 张大鹏
@File : article
@Software: Goland2021.3.1
@Description:
*/

// RandomArticle 生成随机的文章
func (r *Random) RandomArticle() Article {
	index := r.Int(0, len(ArticleList))
	return ArticleList[index]
}
