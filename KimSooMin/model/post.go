package model

type Post struct {
	PostId   int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateAt string `json:"created_at"`
}
