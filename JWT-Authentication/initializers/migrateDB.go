package initializers

import "github.com/PranitRout07/Practice-Golang/JWT-Authentication/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
