package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	config "api.beermenu.com/components/config"
)

var DB *sql.DB

func init() {

	log.Println("Database Package")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Data.DB.User, config.Data.DB.Password, config.Data.DB.Server, config.Data.DB.Port, config.Data.DB.Name)

	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

}
