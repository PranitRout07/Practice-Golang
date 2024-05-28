package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      float64 `json':"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Platform string `json:"platform"`
}

// fake database
var courses []Course

// middleware
func (c *Course) checkCourseNamePresent() bool {
	return c.CourseName == ""
}

func main() {
	//mux engine
	r := mux.NewRouter()

	//seeding

	courses = append(courses, Course{CourseID: "1", CourseName: "AWS Masters", Price: 599.0, Author: &Author{FullName: "Pranit", Platform: "Udemy"}})
	courses = append(courses, Course{CourseID: "2", CourseName: "DevOps Masters", Price: 699.0, Author: &Author{FullName: "Rohan", Platform: "Ingenious"}})
	courses = append(courses, Course{CourseID: "3", CourseName: "K8s Masters", Price: 799.0, Author: &Author{FullName: "Raghav", Platform: "GSG"}})

	//routes
	r.HandleFunc("/", Home)
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetCourseById).Methods("GET")
	r.HandleFunc("/course", CreateCourse).Methods("POST")
	r.HandleFunc("/course/{id}",UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",DeleteCourse).Methods("DELETE")

	fmt.Println("Running at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Study Platform</h1>"))
}
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}
func GetCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No courseid present")

}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Send some data")
	}




	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.checkCourseNamePresent() {
		json.NewEncoder(w).Encode("Enter the course name and then send!")
		return
	}
	random := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(random)
	courseId := strconv.Itoa(rng.Intn(15))
	course.CourseID = courseId
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

}

func UpdateCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for index,course := range courses{
		if course.CourseID==params["id"]{
			//remove the course
			currentCourseCopy := courses
			courses = append(courses[:index],courses[index+1:]...)
			//decode and add the new updated code
			var updatedCourse Course
			json.NewDecoder(r.Body).Decode(&updatedCourse)	
			if updatedCourse.checkCourseNamePresent(){
				courses = currentCourseCopy
				json.NewEncoder(w).Encode("Give the name")
				return
			}
		
			updatedCourse.CourseID = course.CourseID
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode("Updated the course")
			return
		}
	}
	json.NewEncoder(w).Encode("No such Id.Send correct course id")
	return
}

func DeleteCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for index,course := range courses {
		if course.CourseID==params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No such ID.Send correct id")
	return
}