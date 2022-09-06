package query

import (
	"database/sql"
	"fmt"

	"github.com/belljun3395/learngo/DB/structs"
)

func GetPost(conn *sql.DB, paramId string) structs.Board {
	var id int
	var title string
	var detail string
	err := conn.QueryRow("select id, title, detail from board where id =?", paramId).Scan(&id, &title, &detail)
	if err != nil {
		fmt.Println(err)
	}
	var board structs.Board
	board = structs.Board{Web_board_id: id, Web_board_title: title, Web_board_detail: detail}

	return board
}
