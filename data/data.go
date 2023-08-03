package data

import (
	"log"
	"fmt"
	"database/sql"		
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)
var (
	campaignTable = "campaigns"
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

func RegisterCampaign(GuildID string, name string) (string, error) {
	gi, _ := strconv.ParseInt(GuildID, 10, 64)
	if !guildHasAtLeastOneCampaign(gi) {
	registerQ := fmt.Sprintf("INSERT INTO %s (guildId, name) VALUES (?, ?);", campaignTable)

	stmt, _ := db.Prepare(registerQ)
	stmt.Exec(gi, name);
	defer stmt.Close()
	} else {
		return "I am sorry, I can only support one campaign per server at this time.", nil
	}
	return "", nil
} 

func GetTodaysWeatherReport(cId int) (string, error) {
	return "Report", nil
}

func checkerr(err error, q string) {
	if err != nil {
		log.Println(q)
		log.Fatal(err)
	}
}

func runQuery(query string) (*sql.Rows, error) {
	return db.Query(query);
}


func guildHasAtLeastOneCampaign(gi int64) bool {
	var count int
	cq := fmt.Sprintf("SELECT COUNT(*) FROM %s c WHERE c.guildID = %d;", campaignTable, gi)
	err := db.QueryRow(cq).Scan(&count)
	checkerr(err, cq)
	if count > 0 {
		return true;
	} 
	return false
}


