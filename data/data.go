package data

import (
	"log"
	"fmt"
	"database/sql"		
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DateLayout = "2006-01-02"

var (
	db *sql.DB
)

func init() {
	path := os.Getenv("CASTELLAN_DATABASE_PATH")
	db, _ = sql.Open("sqlite3", path)
    // Set maximum idle connections
    db.SetMaxIdleConns(10)

    // Set maximum open connections
    db.SetMaxOpenConns(100)
	loadValidWeatherSystems()
	log.Println("database initlized")
}

func loadValidWeatherSystems() {
	tableq :=  fmt.Sprintf("SELECT id, systemName FROM %s;", weatherSystemTable);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err)
	for rows.Next() {
		var id int
		var name string;
		err = rows.Scan(&id, &name) 
		checkerr(err)
		validWeatherSystems[name] = id;
	}
	log.Println(validWeatherSystems)
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

func logerr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func runQuery(query string, params ...interface{}) (*sql.Rows, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	return stmt.Query(params...);
}



