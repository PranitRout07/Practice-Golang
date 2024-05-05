package main

import (
	"fmt"

	"github.com/PranitRout07/Practice-Golang/CRUD-using-gin-package/controllers"
	internal "github.com/PranitRout07/Practice-Golang/CRUD-using-gin-package/internal/database"
	"github.com/PranitRout07/Practice-Golang/CRUD-using-gin-package/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := internal.InitDB()

	if db == nil {
		fmt.Println("Error while connecting the database")
	}
	dataService := &services.DataService{}
	dataService.InitService(db)

	controller := &controllers.Api{}
	controller.Initialize(r)

	r.Run(":9090")

}
