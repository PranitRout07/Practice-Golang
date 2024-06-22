package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	DB = append(DB, Task{1, "A Todo App", false})
	DB = append(DB, Task{2, "CHI-Project", false})
	DB = append(DB, Task{3, "Use HTMX", false})
	router := chi.NewMux()
	log.Println(middleware.Logger(router))
	router.Get("/", HomeHandler)

	router.Post("/post", PostHandler)
	// router.Get("/tasks", GetTasks)
	router.Delete("/tasks/{id}/delete",DeleteHandler)
	router.Put("/tasks/{id}/update",UpdateStatusHandler)
	fs := http.FileServer(http.Dir("./templates"))
	router.Handle("/templates/*", http.StripPrefix("/templates/", fs))	
	log.Println("Running port 8080...")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})
	ctx["tasks"] = DB
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, ctx)
	if err != nil {
		log.Fatal(err)
	}


}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	log.Println("From get")
	ctx := make(map[string]interface{})
	ctx["tasks"] = DB
	t, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	param := r.FormValue("addtask")
	ctx := make(map[string]interface{})

	val := CheckDuplicateData(param)
	if val{
		t, err := template.ParseFiles("templates/duplicate.html")
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Duplicate data not allwed!!")
		return
	}

	val = CheckEmptyName(param)

	if val{
		t, err := template.ParseFiles("templates/empty.html")
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("empty string not allwed!!")
		return
	}

	length := len(DB) + 1
	DB = append(DB, Task{length, param, false})
	
	// ctx["status"] = "Added a new task successfully!"
	ctx["tasks"] = DB

	t, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(DB, "after adding")

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.Atoi(chi.URLParam(r,"id"))
	if err!=nil{
		log.Fatal(err)
	}

	var index int 
	for i:=0;i<len(DB);i++{
		if DB[i].ID == id {
			index = i 
		}
	}
	
	if len(DB)-1==index{
		DB=DB[0:index]
		log.Println("new delete",DB)
	}else{
		DB = append(DB[0:index],DB[index+1:]...)
	}
	
	ctx := make(map[string]interface{})
	ctx["tasks"] = DB
	t, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(chi.URLParam(r,"id"))
	if err!=nil{
		log.Fatal(err)
	}
	

	for i:=0;i<len(DB);i++{
		if DB[i].ID == id{
			if !DB[i].Status{
				DB[i].Status = true 
				break
			}else{
				DB[i].Status = false
				break
			}
			
		}
	}
	ctx := make(map[string]interface{})
	ctx["tasks"] = DB
	t, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("After status update DB is:",DB)

}


func CheckDuplicateData(params string) bool {
	for i:=0;i<len(DB);i++{
		if DB[i].TaskName == params {
			return true
		}
	}
	return false
}

func CheckEmptyName(params string) bool{
	return params==""
}