package api

import (
	"encoding/json"
	"goRest/api"
	"goRest/db"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOne(id)
	if err != nil {
		api.HandleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		api.HandleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}
