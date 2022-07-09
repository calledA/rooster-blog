package models

type Tag struct {
	Id int
	Name string
	CreatedOn int
	CreatedBy string
	TagClick int
	State int
}

func GetTopicRank() (tag []Tag,err error) {
	if err = DB.Order("tag_click desc").Limit(10).Find(&tag).Error;err != nil {
		return tag,err
	}
	return tag,nil
}