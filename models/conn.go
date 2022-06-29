package models

import (
	"fmt"
	"rooster-blog/pkg/logging"
	"rooster-blog/pkg/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB(conf setting.ServerConfig) {
	var err error

	dbParams := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
	)

	DB,err = gorm.Open("mysql",dbParams)

	if err != nil {
		logging.Fatal("open mysql error",err)
	}

	// 全局禁用表名复数
	DB.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

}