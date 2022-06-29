package setting

import (
	"github.com/go-ini/ini"
	"rooster-blog/pkg/logging"
	"time"
)

var Cfg *ini.File

//服务端配置数据结构
type ServerConfig struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Type         string
	User         string
	Password     string
	Host         string
	DbName       string
	TablePrefix  string
	RedisHost    string
	RedisIndex   string
	PageSize     int
	JwtSecret    string
}

//加载服务端配置
func LoadServerConfig() ServerConfig {

	var err error

	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		logging.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}
	//server配置节点读取
	server, err := Cfg.GetSection("server")
	if err != nil {
		logging.Fatal("Fail to get section 'server': %v", err)
	}
	//app配置节点读取
	app, err := Cfg.GetSection("app")
	if err != nil {
		logging.Fatal("Fail to get section 'app': %v", err)
	}

	//database配置节点读取
	database, err := Cfg.GetSection("database")
	if err != nil {
		logging.Fatal("Fail to get section 'database': %v", err)
	}
	//redis 配置节点读取
	redis, err := Cfg.GetSection("redis")
	if err != nil {
		logging.Fatal("Fail to get section 'redis': %v", err)
	}

	Config := ServerConfig{
		RunMode:      Cfg.Section("").Key("RUN_MODE").MustString("debug"),
		HTTPPort:     server.Key("HTTP_PORT").MustInt(),
		ReadTimeout:  time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second,
		WriteTimeout: time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second,
		Type:         database.Key("TYPE").MustString(""),
		User:         database.Key("USER").MustString(""),
		Password:     database.Key("PASSWORD").MustString(""),
		Host:         database.Key("HOST").MustString(""),
		DbName:       database.Key("NAME").MustString(""),
		TablePrefix:  database.Key("TABLE_PREFIX").MustString(""),
		RedisHost:    redis.Key("HOST").MustString(""),
		RedisIndex:   redis.Key("INDEX").MustString(""),
		PageSize:     app.Key("PAGE_SIZE").MustInt(10),
		JwtSecret:    app.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)"),
	}

	return Config
}
