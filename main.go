package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()
	// Route handles & endpoints
	//4.1
	r.HandleFunc("/api/usersActive", GetUsersActive).Methods("GET")
	//4.2
	r.HandleFunc("/api/usersPostsComments", GetUsersPostsComments).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}

// 4.1
func GetUsersActive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := []Users{}
	json.Unmarshal(getDataFormUrl("https://gorest.co.in/public/v2/users"), &users)
	results := []Users{}
	for _, row := range users {
		if row.Status == "active" {
			results = append(results, row)
		}
	}

	json.NewEncoder(w).Encode(results)
}

// 4.2
func GetUsersPostsComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//users
	users := []Users{}
	json.Unmarshal(getDataFormUrl("https://gorest.co.in/public/v2/users"), &users)

	//Posts
	posts := []Posts{}
	json.Unmarshal(getDataFormUrl("https://gorest.co.in/public/v2/posts"), &posts)

	//comments
	comments := []Commnets{}
	json.Unmarshal(getDataFormUrl("https://gorest.co.in/public/v2/comments"), &comments)

	results := []ResponseUsersPostsComments{}
	for _, user := range users {
		userPosts := []Posts{}
		for _, post := range posts {
			if post.UserId == user.Id {
				postComments := []Commnets{}
				for _, comment := range comments {
					if post.Id == comment.PostId {
						postComments = append(postComments, comment)
					}
				}

				if len(postComments) != 0 {
					post.Commnets = postComments
					post.CommentAmount = len(postComments)
				}
				userPosts = append(userPosts, post)
			}
		}

		responseUsersPostsComment := ResponseUsersPostsComments{
			User: user,
		}
		if len(userPosts) != 0 {
			responseUsersPostsComment.Posts = userPosts
			responseUsersPostsComment.PostAmount = len(userPosts)
		}
		results = append(results, responseUsersPostsComment)
	}

	json.NewEncoder(w).Encode(results)
}

func getDataFormUrl(url string) []byte {
	rows, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rows.Body.Close()

	body, err := ioutil.ReadAll(rows.Body)
	if err != nil {
		panic(err)
	}
	return body
}
