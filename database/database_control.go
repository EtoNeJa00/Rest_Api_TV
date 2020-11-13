package database

import (
	conf "github.com/EtoNeJa00/Rest_Api_TV/configuraton"
	"fmt"
	sql "database/sql"
	_ "github.com/lib/pq"
)

func CreateDB (config conf.DataBase) *sql.DB {
	connStr := fmt.Sprintf(
		"user=%[1]v password=%[2]v host=%[3]v port=%[4]v dbname=%[5]v sslmode=disable", 
		config.User,
		config.Password,
		config.IP,
		config.Port,
		config.DbName)

	db, err := sql.Open("postgres",connStr)
	if err != nil {
		fmt.Print(err)
	}

	return db
}