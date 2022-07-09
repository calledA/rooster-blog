package models

type Article struct {
	Id           int
	TagId        string
	Title        string
	Desc         string
	Content      string
	CreatedOn    string
	CreatedBy    string
	ModifiedOn   string
	ModifiedBy   string
	State        int
	ArticleClick int
	ArticleIp    string
}


func GetArticleRank() (article []Article,err error) {
	if err = DB.Order("article_click desc").Limit(10).Find(&article).Error;err != nil {
		return article,err
	}
	return article,nil
}

func GetClicks() (total []Article,err error) {
	if err := DB.Find(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}
