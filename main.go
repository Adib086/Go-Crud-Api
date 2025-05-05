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
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUsersById).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUsers).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUsers).Methods("DELETE")

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
