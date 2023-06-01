package usecase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"ogs_test/model"
)

func UsersPostsComments() []model.ResponseUsersPostsComments {
	//users
	users := []model.Users{}
	json.Unmarshal(getDataFormUrl("https://gorest.co.in/public/v2/users"), &users)

	//Posts
	posts := []model.Posts{}
	json.Unmarshal(getDataFormUrl("https://gorest.co.in/public/v2/posts"), &posts)

	//comments
	comments := []model.Commnets{}
	json.Unmarshal(getDataFormUrl("https://gorest.co.in/public/v2/comments"), &comments)

	results := []model.ResponseUsersPostsComments{}
	for _, user := range users {
		userPosts := []model.Posts{}
		for _, post := range posts {
			if post.UserId == user.Id {
				postComments := []model.Commnets{}
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

		responseUsersPostsComment := model.ResponseUsersPostsComments{
			User: user,
		}
		if len(userPosts) != 0 {
			responseUsersPostsComment.Posts = userPosts
			responseUsersPostsComment.PostAmount = len(userPosts)
		}
		results = append(results, responseUsersPostsComment)
	}
	return results
}

func UserActive() []model.Users {
	users := []model.Users{}
	json.Unmarshal(getDataFormUrl("https://gorest.co.in/public/v2/users"), &users)
	results := []model.Users{}
	for _, row := range users {
		if row.Status == "active" {
			results = append(results, row)
		}
	}
	return results
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
