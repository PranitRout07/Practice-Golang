package routes

import (
	"github.com/PranitRout07/Practice-Golang/login-logout/controllers"
	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	router := chi.NewMux()

	router.Get("/", controllers.Home)
	router.Get("/loginform",controllers.LoginForm)
	router.Get("/registerform",controllers.RegisterForm)
	router.Post("/register",controllers.Register)
	router.Post("/login",controllers.Login)
	router.Post("/logout",controllers.Logout)
	router.Get("/content",controllers.GetContent)
	return router

}
