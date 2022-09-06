package query

import (
	"database/sql"
	"fmt"

	"github.com/belljun3395/learngo/DB/structs"
)

func PostWrite(conn *sql.DB, write structs.Write) bool {
	_, err := conn.Exec("INSERT INTO board (title, detail) VALUES (?, ?) ", write.Title, write.Detail)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
