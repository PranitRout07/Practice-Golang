package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/initializers"
	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/middlewares"
	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/models"
	"github.com/go-chi/chi/v5"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	ctx := make(map[string]string)
	ctx["post"] = "posts"
	err = t.Execute(w, ctx)
	if err != nil {
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
	for rows.Next() {
		var data models.Posts
		err := rows.Scan(&data.Id, &data.Title)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, data)
	}

	t, err := template.ParseFiles("templates/pages/posts.html")
	if err != nil {
		log.Fatal(err)
	}
	ctx := make(map[string]interface{})
	ctx["posts"] = posts
	err = t.Execute(w, ctx)
	if err != nil {
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
	for rows.Next() {
		var data models.Products
		err := rows.Scan(&data.Id, &data.ProductName, &data.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, data)
	}

	t, err := template.ParseFiles("templates/pages/products.html")
	if err != nil {
		log.Fatal(err)
	}
	ctx := make(map[string]interface{})
	ctx["products"] = products
	err = t.Execute(w, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(products)
}

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from detail handler")

	// post := r.Context().Value(middlewares.PostsKey)
	post, ok := r.Context().Value(middlewares.PostsKey).(models.Posts)
	if !ok {
		log.Println("No post found in context")
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	log.Println(post)
	t, err := template.ParseFiles("templates/pages/details.html")
	if err != nil {
		log.Fatal(err)
	}
	ctx := make(map[string]interface{})
	ctx["posts"] = post
	err = t.Execute(w, ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func PostArticles(w http.ResponseWriter, r *http.Request) {

	article := r.FormValue("article")

	if middlewares.CheckDuplicateDatas(article){
		
		ctx := make(map[string]interface{})
		ctx["result"] = "Data with same name is already present in the database. Please Use another name!"
		t, _ := template.ParseFiles("templates/pages/responseForPost.html")
		err := t.Execute(w, ctx)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	sqlQuery := fmt.Sprintf("INSERT INTO posts(title) VALUES ('%s');", article)

	fmt.Println(sqlQuery)

	// Execute the SQL query
	res, err := initializers.DBConnection.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil {
		ctx := make(map[string]interface{})
		ctx["result"] = "Successfully added!"
		t, _ := template.ParseFiles("templates/pages/responseForPost.html")
		err := t.Execute(w, ctx)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func DeleteArticles(w http.ResponseWriter, r *http.Request) {
	log.Println("This is from delete handler.")

	id := chi.URLParam(r, "id")
	log.Println("Delete id", id)
	ID, _ := strconv.Atoi(id)

	sqlQuery := "DELETE FROM posts WHERE id=$1;"
	res, err := initializers.DBConnection.Exec(sqlQuery, ID)

	log.Println("Print delete result", res)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		t, _ := template.ParseFiles("templates/pages/deletePosts.html")
		err := t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func UpdateData(w http.ResponseWriter, r *http.Request){

	log.Println("This is from update handler.")
	article := r.FormValue("article")
	id := chi.URLParam(r, "id")
	log.Println("Update id", id)
	ID, _ := strconv.Atoi(id)

	sqlQuery := fmt.Sprintf("UPDATE posts SET title = '%s' WHERE id=$1;",article)
	res, err := initializers.DBConnection.Exec(sqlQuery, ID)

	log.Println("Print update result", res)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		t, _ := template.ParseFiles("templates/pages/updatePosts.html")
		err := t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}

	}


}


