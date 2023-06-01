package model

type ResponseUsersPostsComments struct {
	User       Users
	PostAmount int     `json:"post_amount"`
	Posts      []Posts `json:"posts"`
}
