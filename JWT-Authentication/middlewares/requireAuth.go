package middlewares

import (
	"fmt"
	"github.com/PranitRout07/Practice-Golang/JWT-Authentication/initializers"
	"github.com/PranitRout07/Practice-Golang/JWT-Authentication/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	//Get the cookie

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Println("Error: No JWT token found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	log.Println(tokenString)

	//validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Println("Error: No JWT token found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//check expiry of jwt
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Println("Error: No JWT token found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			return
		}

		fmt.Println(claims["foo"], claims["nbf"])

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			log.Println("Error: No JWT token found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			return
		}

		//set to user

		c.Set("user", user)

		c.Next() //helps in proceeding to execute next middleware or handler

	} else {
		log.Println("Error: No JWT token found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	//Now we have th etoken . We have to find the user of the token from db

}

func RemoveCookies(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Logged out successfully!",
	})
}
