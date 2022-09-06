package db

import (

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	"github.com/kssumin/golang_board/model"

)

func WritePost(title string, details string,conn *sql.DB) bool {
	_, err := conn.Query("INSERT INTO post(title, content, createAt) VALUES (?, ?, ?)",
		title, details, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		panic(err)
	}

	return true
}

func GetPostById(conn *sql.DB, paramId string) model.Post {
	var id int
	var title string
	var content string

	err := conn.QueryRow("select postId, title, content from board where id =?", paramId).Scan(&id, &title, &content)
	if err != nil {
		fmt.Println(err)
	}
	
	var post model.Post
	post = model.Post{PostId: id, Title: title, Content: content}

	return post
}

func GetPosts(conn *sql.DB) [] model.Post {
	rows,_:= conn.Query("select id, title, content from post")
	var posts []model.Post
	for rows.Next() {
		var id int
		var title string
		var content string
		var post model.Post
		
		rows.Scan(&id, &title, &content)
		post = model.Post{PostId: id, Title: title, Content: content}
		posts = append(posts, post)
	}
	return posts
}
