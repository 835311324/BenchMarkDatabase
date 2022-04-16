package main

import (
	"BenchMarkDataBase/conf"
	"BenchMarkDataBase/router"
)

func main() {
	conf.InitConf("./conf/app.ini")
	r := router.NewRouter()
	_ = r.Run()
}
