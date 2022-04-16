package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func InitDatabase(dsn string) {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
}
