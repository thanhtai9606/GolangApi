package config

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (db *sql.DB, err error) {

	db, err = sql.Open("mssql", "server=10.20.46.23;user id=sa;password=@Sql2016;database=University")
	return
}
