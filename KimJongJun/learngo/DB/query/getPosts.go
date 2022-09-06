package query

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/belljun3395/learngo/DB/structs"
)

func GetPosts(conn *sql.DB) []structs.Board {
	rows, err := conn.Query("select id, title, detail from board")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var boards []structs.Board
	for rows.Next() {
		var id int
		var title string
		var detail string
		var board structs.Board
		rows.Scan(&id, &title, &detail)
		board = structs.Board{Web_board_id: id, Web_board_title: title, Web_board_detail: detail}
		boards = append(boards, board)
	}
	return boards
}
