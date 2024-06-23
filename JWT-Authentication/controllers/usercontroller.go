package controllers

import (

	"github.com/PranitRout07/Practice-Golang/JWT-Authentication/initializers"
	"github.com/PranitRout07/Practice-Golang/JWT-Authentication/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//Get email and password
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	//Create the user and store to db

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//Give success response if user created successfully

	c.JSON(http.StatusOK, gin.H{
		"msg": "Successfully created an user",
	})
}

func Login(c *gin.Context) {
	//get the user login details
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// search the details in the database
	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to retrieve email",
		})
		return
	}

	//compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	//if passowrd check is successful , then create a jwt token.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create JWT token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode) //this is used to stop csrf attacks.
	c.SetCookie("Authorization",tokenString,3600*24*30,"","",false,true) //this sets the cookie and give access the user to see the contents

}

func Validate(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg": "Logged in successfully!",
	})
}

func Logout(c *gin.Context){

}