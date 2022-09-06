package DB

import (
	"database/sql"
)

// func Setup() sql.DB {
// 	conn, err := sql.Open("mysql", "root:kim@2737169@tcp(localhost:3306)/go_crud")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	return *conn
// }

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "kim@2737169"
	dbName := "go_crud"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
