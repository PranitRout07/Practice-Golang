package main

import (
	"github.com/PranitRout07/Practice-Golang/JWT-Authentication/controllers"
	"github.com/PranitRout07/Practice-Golang/JWT-Authentication/initializers"
	"github.com/PranitRout07/Practice-Golang/JWT-Authentication/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.MigrateDB()
}

func main() {
	router := gin.Default()
	router.POST("/signup", controllers.SignUp)
	router.POST("/login",controllers.Login)
	router.GET("/validate",middlewares.RequireAuth,controllers.Validate)
	router.POST("/logout",middlewares.RemoveCookies,controllers.Logout)
	log.Fatal(http.ListenAndServe(":4000", router))

}
