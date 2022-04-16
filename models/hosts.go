package models

import "BenchMarkDataBase/database"

type Host struct {
	id uint
}

func (h *Host) AddOne() (int64, error) {
	return database.ModifyDB("insert into hosts (id ) values (?)", 100)
}
