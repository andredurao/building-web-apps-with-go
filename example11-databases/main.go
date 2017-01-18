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
		// Iterate through the records
		for rows.Next(){
			var row Book
			// Scan title and author on row attribute references
			err = rows.Scan(&row.Title, &row.Author)
			if err != nil{
				panic(err)
			}
			// Append the row on result map
			result = append(result, row)
			
			// TODO: Iterate on the map at the view and render the results on a template
			for i, value := range result {
				fmt.Fprintf(rw, "<li>%d: '%s' by '%s'</li>\n", i, value.Title, value.Author)
			}			
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
