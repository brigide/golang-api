package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	/* Users Routes */
	r.HandleFunc("/users", controllers.create).Methods("POST")

	return r
}
