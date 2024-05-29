package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	

	// define handlers
	http.HandleFunc("/", GetDetails)
	http.HandleFunc("/add-film/", AddFilms)

	fmt.Println("Running in port 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func GetDetails(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "Avengers", Director: "Anthony Russo"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}



func AddFilms(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}