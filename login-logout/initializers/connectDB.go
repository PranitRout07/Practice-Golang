package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DBConnection *sql.DB

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",os.Getenv("host"),os.Getenv("user"),os.Getenv("password"),os.Getenv("dbname"),os.Getenv("port"))
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	DBConnection = db
	log.Println("Connected to database...")
	return DBConnection
}
