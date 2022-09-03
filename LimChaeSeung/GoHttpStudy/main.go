package main

import (
	"GoHttpStudy/db"
	"GoHttpStudy/router"
	"database/sql"
	"net/http"
)

func main() {
	var err error
	db.DB, err = sql.Open("mysql", "loopy:loopy@tcp(loopy-lim.com:3307)/loopy")
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	r := router.NewRouters()
	go r.Run()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":3000", nil)

}
