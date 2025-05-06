package main

import (
	"GoCrudApi/types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var users []types.User
var courses []types.Course

func getUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}
func deleteUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range users {
		if item.Id == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(users)
}

func getUsersById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range users {
		if item.Id == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}
func createUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user types.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	user.Id = strconv.Itoa(rand.Intn(10000000))
	users = append(users, user)
	json.NewEncoder(writer).Encode(user)

}

func updateUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range users {
		if item.Id == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user types.User
			_ = json.NewDecoder(request.Body).Decode(&user)
			user.Id = params["id"]
			users = append(users, user)
			json.NewEncoder(writer).Encode(user)
			return
		}
	}
}

func getCourses(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(courses)
}
func deleteCourse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(users)
}

func getCourseById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range courses {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}
func createCourse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var course types.Course
	_ = json.NewDecoder(request.Body).Decode(&course)
	course.ID = strconv.Itoa(rand.Intn(10000000))
	courses = append(courses, course)
	json.NewEncoder(writer).Encode(course)

}

func updateCourse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course types.Course
			_ = json.NewDecoder(request.Body).Decode(&course)
			course.ID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(writer).Encode(course)
			return
		}
	}
}

func main() {
	fmt.Println("Hello World")
	r := mux.NewRouter()
	users = append(users, types.User{
		Id:   "12",
		Name: "Adib",
		Age:  22,
		Courses: []types.Course{
			{
				ID:    "1",
				Title: "Bangla",
			},
		},
	})

	courses = append(courses, types.Course{
		ID:    "2",
		Title: "English",
	})
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUsersById).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUsers).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUsers).Methods("DELETE")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
