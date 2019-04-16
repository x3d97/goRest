package main

import (
	"net/http"

	api "goRest/api/users"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", api.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/api/users", api.AddUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", api.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
