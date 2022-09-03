package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB

type Board struct {
	WebBoardId     int    `json:"web_board_id,omitempty"`
	WebBoardTitle  string `json:"web_board_title,omitempty"`
	WebBoardDetail string `json:"web_board_detail"`
	DateTime       string `json:"date_time,omitempty"`
}

func GetPostsByDB() []Board {
	rows, err := DB.Query("SELECT web_board_id, web_board_title, web_board_details, date_time FROM web_board")
	if err != nil {
		fmt.Println(err)
		return []Board{}
	}
	defer rows.Close()

	boards := make([]Board, 0)
	for rows.Next() {
		var board Board
		err = rows.Scan(&board.WebBoardId, &board.WebBoardTitle, &board.WebBoardDetail, &board.DateTime)
		boards = append(boards, board)
	}

	return boards
}

func GetPostByDB(id int) Board {
	board := Board{}
	err := DB.QueryRow("SELECT web_board_id, web_board_title, web_board_details, date_time FROM web_board WHERE web_board_id=?",
		id).Scan(&board.WebBoardId, &board.WebBoardTitle, &board.WebBoardDetail, &board.DateTime)
	if err != nil {
		fmt.Println(err)
		return Board{}
	}

	return board
}

func WritePostToDB(title string, details string) bool {
	_, err := DB.Exec("INSERT INTO web_board(web_board_title, web_board_details, date_time) VALUES (?, ?, ?)",
		title, details, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		panic(err)
		return false
	}

	return true
}
