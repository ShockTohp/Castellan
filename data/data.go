package data

import (
	"log"
	"database/sql"		
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("sqlite3", "./data/castellan.db")
	log.Println("database initlized")
    // Set maximum idle connections
    db.SetMaxIdleConns(10)

    // Set maximum open connections
    db.SetMaxOpenConns(100)
}

func CloseDatabase() {
	if db != nil {
	db.Close()
	}

}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runQuery(query string) (*sql.Rows, error) {
	return db.Query(query);
}



