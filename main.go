package main

import (
	"log"
	"net/http"
	"ogs_test/controller"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	//4.1
	r.HandleFunc("/api/usersActive", controller.GetUsersActive).Methods("GET")
	//4.2
	r.HandleFunc("/api/usersPostsComments", controller.GetUsersPostsComments).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
