package controller

import (
	"encoding/json"
	"net/http"
	"ogs_test/usecase"
)

// 4.1
func GetUsersActive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usecase.UserActive())
}

// 4.2
func GetUsersPostsComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usecase.UsersPostsComments())
}
