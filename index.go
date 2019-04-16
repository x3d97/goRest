package main

import (
	"fmt"
	"net/http"

	"github.com/x3d97/goRest/config"

	"github.com/caarlos0/env"
	api "github.com/x3d97/goRest/api/users"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Config{}
	env.Parse(&cfg)
	r := mux.NewRouter()
	r.HandleFunc("/", mainPage)
	r.HandleFunc("/api/users", api.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/api/users", api.AddUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", api.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":"+cfg.Port, r)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Kuk")
}
