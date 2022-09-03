package router

import (
	"GoHttpStudy/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type writePostStruct struct {
	title   string
	details string
}

func GetPosts(c *gin.Context) {
	boards := db.GetPostsByDB()
	fmt.Println(boards)
	c.JSON(200, boards)
}

func GetPost(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	board := db.GetPostByDB(int(id))
	fmt.Println(board)
	c.JSON(200, board)
}

func WritePost(c *gin.Context) {
	c.Request.ParseForm()
	data := writePostStruct{c.PostForm("title"), c.PostForm("details")}

	db.WritePostToDB(data.title, data.details)
	c.Redirect(301, "http://localhost:3000/home.html")
}
