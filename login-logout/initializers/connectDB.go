package initializers

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DBConnection *sql.DB

func ConnectDB() *sql.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5440 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	DBConnection = db
	log.Println("Connected to database...")
	return DBConnection
}
