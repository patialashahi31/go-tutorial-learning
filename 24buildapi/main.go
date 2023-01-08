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

// model for course - file

type Course struct {
	CourseId  string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
} 

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake db

var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool{
	return c.CourseName == ""
}

func main() {
	fmt.Println("APIs")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2",CourseName: "React",CoursePrice: 123, Author: &Author{
		Fullname: "Tej", Website: "hsh.com",
	}})
	courses = append(courses, Course{CourseId: "3",CourseName: "Python",CoursePrice: 23, Author: &Author{
		Fullname: "Kk", Website: "c.com",
	}})

	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("DELETE")
	r.HandleFunc("/course",createOneCourse).Methods("POST")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000",r))
}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1> Hello APIs </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses ")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("get one course")
	w.Header().Set("Content-Type","application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses and find matching id
	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with give id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course ")
	w.Header().Set("Content-Type","application/json")

	// What if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	// what if data = {}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data")
		return
	}

	// generate a unique id , string
	// append courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course ")
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index,course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: Handle when id is not matched
}

func deleteOneCourse (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course ")
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index,course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			json.NewEncoder(w).Encode("Course removed")
			break
		}
	}

	// TODO: Handle when id is not matched
}