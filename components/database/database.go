package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {

	log.Println("Database Package")

	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:8889)/%s", "root", "root", "WineMenu")

	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

}
