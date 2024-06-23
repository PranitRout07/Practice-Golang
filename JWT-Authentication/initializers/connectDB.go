package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5440 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to database.")

	}

	DB = db

	log.Println("Connected to database successfully")
}
