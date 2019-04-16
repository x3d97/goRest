package api

import (
	"encoding/json"
	"goRest/api"
	"goRest/db"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAll()
	if err != nil {
		api.HandleError(err, "Failed to load database items: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		api.HandleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}
