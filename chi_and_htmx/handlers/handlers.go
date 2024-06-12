package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/initializers"
	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	ctx := make(map[string]string)
	ctx["post"] = "posts"
	err = t.Execute(w, ctx)
	if err!=nil{
		log.Fatal(err)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	var posts []models.Posts
	sql_querry := "select * from posts" //create a table named post with all its fields

	rows, err := initializers.DBConnection.Query(sql_querry)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next(){
		var data models.Posts
		err:= rows.Scan(&data.Id,&data.Title)
		if err!=nil{
			log.Fatal(err)
		}
		posts = append(posts, data)
	}

	t, err := template.ParseFiles("templates/pages/posts.html")
	if err!=nil{
		log.Fatal(err)
	}
	ctx := make(map[string]interface{})
	ctx["posts"] = posts
	err = t.Execute(w, ctx)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println(posts)
	
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Products
	sql_querry := "select * from products" //create a table named post with all its fields

	rows, err := initializers.DBConnection.Query(sql_querry)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next(){
		var data models.Products
		err:= rows.Scan(&data.Id,&data.ProductName,&data.Price)
		if err!=nil{
			log.Fatal(err)
		}
		products = append(products, data)
	}

	t, err := template.ParseFiles("templates/pages/products.html")
	if err!=nil{
		log.Fatal(err)
	}
	ctx := make(map[string]interface{})
	ctx["products"] = products
	err = t.Execute(w, ctx)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println(products)
}
