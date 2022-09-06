package main

import (
	"database/sql"
	"fmt"
	"time"

	db "github.com/belljun3395/learngo/DB"
	"github.com/belljun3395/learngo/DB/query"
	"github.com/belljun3395/learngo/DB/structs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func setupRouter(conn *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"POST", "GET"},
			MaxAge:       12 * time.Hour,
		}))

	r.GET("/api/getPosts", func(c *gin.Context) {
		boards := query.GetPosts(conn)
		for _, board := range boards {
			fmt.Printf("id : %d, title : %s, detail : %s \n", board.Web_board_id, board.Web_board_title, board.Web_board_detail)
		}
		c.JSON(200, boards)
	})

	r.GET("/api/getPost", func(c *gin.Context) {
		id := c.Query("id")
		board := query.GetPost(conn, id)
		fmt.Printf("id : %d, title : %s, detail : %s \n", board.Web_board_id, board.Web_board_title, board.Web_board_detail)
		c.JSON(200, board)
	})

	r.POST("/api/writePost", func(c *gin.Context) {
		title := c.PostForm("title")
		detail := c.PostForm("detail")
		write := structs.Write{Title: title, Detail: detail}
		fmt.Println(write)
		query.PostWrite(conn, write)
		c.JSON(200, write)
	})

	return r
}

func main() {
	conn := db.DbConn()
	defer conn.Close()

	r := setupRouter(conn)
	r.Run(":8080")

}
