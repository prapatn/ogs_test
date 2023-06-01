package model

type Posts struct {
	Id            int        `json:"id"`
	UserId        int        `json:"user_id"`
	Title         string     `json:"title"`
	Body          string     `json:"body"`
	CommentAmount int        `json:"comment_amount"`
	Commnets      []Commnets `json:"comments"`
}
