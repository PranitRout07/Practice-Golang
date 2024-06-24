package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/PranitRout07/Practice-Golang/login-logout/initializers"
	"github.com/PranitRout07/Practice-Golang/login-logout/middlewares"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func LoginForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	log.Println(email, password)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Fatal(err)
	}

	sqlQuery := fmt.Sprintf("INSERT INTO person(email,password) VALUES ('%s','%s');", email, string(hash))

	fmt.Println(sqlQuery)

	// Execute the SQL query
	res, err := initializers.DBConnection.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil {
		ctx := make(map[string]interface{})
		ctx["result"] = "Successfully added!"
		t, _ := template.ParseFiles("templates/responseAfterRegister.html")
		err := t.Execute(w, ctx)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	//check whether the email present in the db or not

	var count int
	sqlQuery := fmt.Sprintf("SELECT COUNT(*) FROM person WHERE email = '%s';", email)
	err := initializers.DBConnection.QueryRow(sqlQuery).Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		log.Println("No such email exists.")
	}

	//compare the both password

	var passwordFromDB string
	var idFromDB int64
	query := "SELECT password, id FROM person WHERE email = $1"
	err = initializers.DBConnection.QueryRow(query, email).Scan(&passwordFromDB, &idFromDB)


	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(passwordFromDB,password)
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(password))
	if err != nil {
		log.Println("Incorrect password")
	}

	//if passowrd check is successful , then create a jwt token.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": idFromDB,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err!=nil{
		log.Println("Error while creating the jwt token",err)
	}
	log.Println("TOKEN:",tokenString)
	http.SetCookie(w,&http.Cookie{
		Name:"Authorization",
		Value: tokenString,
		Expires: time.Now().Add(time.Hour * 24 * 30) , 
		SameSite: http.SameSiteLaxMode,
	})


	t, err := template.ParseFiles("templates/responseForLogin.html")
	if err!=nil{
		log.Fatal(err)
	}
	err = t.Execute(w,nil)
	if err!=nil{
		log.Fatal(err)
	}

}


func Logout(w http.ResponseWriter, r *http.Request){

	http.SetCookie(w,&http.Cookie{
		Name:"Authorization",
		Value: "",
		Expires: time.Now().Add(0) , 
		SameSite: http.SameSiteLaxMode,
	})

	t,err := template.ParseFiles("templates/logout.html")
	if err!=nil{
		log.Fatal(err)
	}
	err = t.Execute(w,nil)
	if err!=nil{
		log.Fatal(err)
	}
}

func GetContent(w http.ResponseWriter,r *http.Request){
	val := middlewares.Validate(w,r)
	if !val{
		log.Println("You are not authorized!Please login to continue seeing the content.")
		w.Write([]byte("You are not authorized!Please login to continue seeing the content."))
		return
	}

	t,err := template.ParseFiles("templates/showContent.html")
	if err!=nil{
		log.Fatal(err)
	}
	err = t.Execute(w,nil)
	if err!=nil{
		log.Fatal(err)
	}
}