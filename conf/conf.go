package conf

import (
	"BenchMarkDataBase/database"
	"BenchMarkDataBase/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

var (
	mysqlHost     string
	mysqlPort     string
	mysqlUser     string
	mysqlPassword string
	mysqlDatabase string

	ginMode string

	logDir string
)

func InitConf(file string) {
	cfg, _ := ini.Load(file)

	//log
	logDir = cfg.Section("log").Key("logdir").String()
	utils.CreatLog(logDir)

	// mysql
	mysqlHost = cfg.Section("mysql").Key("host").String()
	mysqlPort = cfg.Section("mysql").Key("port").String()
	mysqlUser = cfg.Section("mysql").Key("user").String()
	mysqlPassword = cfg.Section("mysql").Key("password").String()
	mysqlDatabase = cfg.Section("mysql").Key("database").String()
	database.InitDatabase(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase))

	// gin
	ginMode = cfg.Section("gin").Key("mode").String()
	gin.SetMode(ginMode)

	utils.AppLog.Info("init conf success")
}
