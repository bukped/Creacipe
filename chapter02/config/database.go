package config

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/creacipe_db?parseTime=true")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	DB = db
	log.Println("database connected")
}