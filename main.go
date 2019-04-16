package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", top)
	http.HandleFunc("/article", article)
	http.ListenAndServe(":8080", nil)
}

func top(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(db:3306)/hoge")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select id, name from users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err.Error())
		}
		fmt.Fprintf(w, "id: %d name: %s\n", id, name)
	}
}

func article(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Article !\n")
}
