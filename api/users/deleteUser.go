package api

import (
	"github.com/x3d97/goRest/api"
	"github.com/x3d97/goRest/db"

	"net/http"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	if err := db.Remove(id); err != nil {
		api.HandleError(err, "Failed to remove item: %v", w)
		return
	}

	w.Write([]byte("OK"))
}
