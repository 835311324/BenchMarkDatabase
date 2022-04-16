package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Weblog *logrus.Logger
var AppLog *logrus.Logger

func CreatLog(logdir string) {
	al := logdir + "/" + "app.log"
	wl := logdir + "/" + "web.log"
	AppLog = initLog(al)
	Weblog = initLog(wl)
}

func initLog(logpath string) *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	var f *os.File
	var err error
	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		f, err = os.Create(logpath)
	} else {
		f, err = os.OpenFile(logpath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	if err != nil {
		fmt.Println("open %s file failed", logpath)
	}

	log.Out = f
	log.Level = logrus.InfoLevel
	return log
}
