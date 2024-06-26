package routes

import (
	"github.com/PranitRout07/Practice-Golang/login-logout/controllers"
	"github.com/go-chi/chi/v5"
)





func Routes() *chi.Mux {
	router := chi.NewMux()
	tempDetailsHandler := &controllers.TempDetail{}

	router.Get("/", controllers.Home)
	router.Get("/loginform", controllers.LoginForm)
	router.Get("/registerform", controllers.RegisterForm)
	router.Post("/register", tempDetailsHandler.Register)
	router.Post("/login", controllers.Login)
	router.Post("/logout", controllers.Logout)
	router.Get("/content", controllers.GetContent)
	router.Post("/otp",tempDetailsHandler.RegisterAfterOTPConfirmation)
	router.Get("/check-otp", controllers.CheckOTPTime)
	router.Post("/resendotp",tempDetailsHandler.Resend)
	return router

}
