package models

import (
	"BenchMarkDataBase/database"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestHosts_AddOne(t *testing.T) {
	database.InitDatabase("hal:hal123@tcp(192.168.240.143:3306)/benchmark?charset=utf8mb4&parseTime=True")
	h := &Host{id: 100}
	row, ret := h.AddOne()
	t.Logf("rows:%d,ret:%s", row, ret.Error())
}
