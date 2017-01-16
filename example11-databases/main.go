package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	db := NewDB()
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", ShowBooks(db))
}

func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {	
		rows, err := db.Query("SELECT title, author FROM books")
		if err != nil {
			panic(err)
		}
		result := []Book{}
		for rows.Next(){
			var row Book
			err = rows.Scan(&row.Title, &row.Author)
			if err != nil{
				panic(err)
			}
			result = append(result, row)
		}
	})
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books(title TEXT, author TEXT)")
	if err != nil {
		panic(err)
	}

	return db
}
