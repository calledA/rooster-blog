package models

import (
	// "fmt"
	"time"
)

type Click struct {
	Id          int
	TrticleId   int
	TrticleName string
	Topic       string
	ClickTime   string
	ClickBy     string
}
func GetClicks() (int64, error) {
	var count int64
	if err := DB.Model(&Click{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func CurrentDateTime() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
}
