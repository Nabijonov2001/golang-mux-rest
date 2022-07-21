package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Datalar uchun model
type Course struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Job      string `json:"job"`
}

// Ma'lumotlar ba'zasi
var courses []Course = []Course{
	{Id: "1", Name: "Node.js microservices", Price: 30.5, Author: &Author{Fullname: "Fazliddin Nabijonov", Job: "Web Developer"}},
	{Id: "2", Name: "Golang basics", Price: 20.4, Author: &Author{Fullname: "Fazliddin Nabijonov", Job: "Web Developer"}},
}

func main() {

	fmt.Println("Go dasturlash tilida CRUD bilan ishlash!")

	// Routni ishga tushirish
	r := mux.NewRouter()

	r.HandleFunc("/api/courses", getAll).Methods("GET")
	r.HandleFunc("/api/courses/{id}", getOne).Methods("GET")
	r.HandleFunc("/api/courses/create", create).Methods("POST")
	r.HandleFunc("/api/courses/{id}", update).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", delete).Methods("DELETE")

	// Portni ulash
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Barcha kurslarni olish
func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Bitta kursni o'zini olish
func getOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("There is no course with this id")
}

// Yangi kurs yaratish
func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	course.Id = strconv.Itoa(len(courses) + 1)
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

// Kursni update qilish
func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var updatedCourse Course
	json.NewDecoder(r.Body).Decode(&updatedCourse)

	for index, course := range courses {
		if course.Id == params["id"] {
			updatedCourse.Id = course.Id
			courses = append(courses[:index], courses[index+1:]...)
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("There is no course with this id")
}

// Kursni o'chirish
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("There is no course with this id")
}
