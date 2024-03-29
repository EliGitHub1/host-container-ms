package dataBase

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const DBPath string = "C:\\Users\\eliav\\Desktop\\new\\host-container-ms\\aqua.db"
const DriverName string = "sqlite3"

func GetDataBase() *sql.DB {
	database, err := sql.Open(DriverName, DBPath)
	if err != nil {
		return nil
	}
	return database
}
