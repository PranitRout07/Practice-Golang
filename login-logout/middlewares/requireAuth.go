package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Validate(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("Authorization")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("Can not get the cookie")
			w.WriteHeader(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusBadGateway)
		return false
	}

	tokenString := cookie.Value
	var claims jwt.MapClaims
	tkn, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Println("Error: No JWT token found")
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}
