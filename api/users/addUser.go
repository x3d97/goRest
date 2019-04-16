package api

import (
	"goRest/api"
	"goRest/db"
	"net/http"
	"strconv"
)

func AddUser(w http.ResponseWriter, req *http.Request) {
	ID := req.FormValue("id")
	name := req.FormValue("name")
	email := req.FormValue("email")
	phone, err := strconv.Atoi(req.FormValue("phone"))

	if err != nil {
		api.HandleError(err, "Failed to parse input data: %v", w)
		return
	}

	// name, err := strconv.Atoi(valueStr)
	// if err != nil {
	// 	api.HandleError(err, "Failed to parse input data: %v", w)
	// 	return
	// }

	user := db.User{ID: ID, Name: name, Email: email, Phone: phone}
	db.Save(user)

	// if err != nil {
	// 	handleError(err, "Failed to save data: %v", w)
	// 	return
	// }

	w.Write([]byte("OK"))
}
