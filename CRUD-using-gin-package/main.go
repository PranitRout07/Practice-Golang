package main

import (
	"github.com/PranitRout07/Practice-Golang/CRUD-using-gin-package/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controllers.Initialize(r)

}
