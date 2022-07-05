package models

import (
	"fmt"
	"rooster-blog/pkg/logging"
	"rooster-blog/pkg/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// "github.com/go-ini/ini"

)

var (
	DB *gorm.DB
	// Cfg *ini.File
)

func init() {
	var err error
	conf, err := setting.Cfg.GetSection("database")
	if err != err {
		logging.Fatal(2, "Fail to get section 'database': %v", err)
	}

	user := conf.Key("USER").MustString("")
	password := conf.Key("PASSWORD").MustString("")
	host := conf.Key("HOST").MustString("")
	dbName := conf.Key("NAME").MustString("")
	tablePrefix := conf.Key("TABLE_PREFIX").String()

	dbParams := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	)

	DB, err = gorm.Open("mysql", dbParams)

	if err != nil {
		logging.Fatal("open mysql error", err)
	}

	// 全局禁用表名复数
	DB.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

}
