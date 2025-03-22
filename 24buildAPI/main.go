package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Models --diff file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// database
var courses []Course

// middlewares or helper --diff file
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	courses = append(courses, Course{CourseId: "3", CourseName: "Cohort 2.0", CoursePrice: 2999, Author: &Author{Fullname: "Harkirat", Website: "100x.dev.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "JS basics", CoursePrice: 299, Author: &Author{Fullname: "Harkirat", Website: "100x.dev.com"}})

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCoursebyId).Methods("GET")
	r.HandleFunc("/course", addNewCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers --diff file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCoursebyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with provided ID.")
	return
}

func addNewCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Body")
		return
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Please provide some data.")
		return
	}

	for _, cor := range courses {
		if cor.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists.")
			return
		}
	}

	rand.Seed(time.Now().UnixNano()) //to generate random number
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //id

	//we will loop use id and remove course make changes add it back by the id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //removing

			var newcourse Course
			json.NewDecoder(r.Body).Decode(&newcourse) //got the updated data from req
			newcourse.CourseId = params["id"]          //as id wont be their in updated data putting it in place

			courses = append(courses, newcourse) //adding updated ones back
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course with this id found.")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //id

	//we will loop use id and remove course make changes add it back by the id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //removing
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course with this id found.")
	return
}
