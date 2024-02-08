package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func ConnectDB() error {
	log.Println("ConnectDB")

	if Db != nil {
		return nil
	}

	var err error
	Db, err = sql.Open("sqlite3", "mydb.db")
	if err != nil {
		return err
	}

	err = Db.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to bd successfully")

	return nil
}

func CloseDB() {

}
