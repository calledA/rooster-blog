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


func CurrentDateTime() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
}
