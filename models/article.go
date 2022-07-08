package models

import "fmt"

type Article struct {
	Id           int
	TagId        int
	Title        string
	Desc         string
	Content      string
	CreatedOn    string
	CreatedBy    string
	ModifiedOn   string
	ModifiedBy   string
	state        int
	ArticleClick int64
	ArticleIp    string
}

func GetArticleRank() {
	var article []Article
	err := DB.Order("article_click desc").Limit(10).Find(&article).Error
	if err != nil {
		return
	}
	fmt.Println(article)
	// return article,err
}
