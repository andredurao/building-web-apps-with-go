package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"path"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	db := NewDB()
	log.Println("Listening on :8080")
	http.ListenAndServe(":8090", ShowBooks(db))
}

func GetBooksMap(db *sql.DB) []Book {
	rows, err := db.Query("SELECT title, author FROM books")
	if err != nil {
		panic(err)
	}
	result := []Book{}
	// Iterate through the records
	for rows.Next() {
		var row Book
		// Scan title and author on row attribute references
		err = rows.Scan(&row.Title, &row.Author)
		if err != nil {
			panic(err)
		}
		// Append the row on result map
		result = append(result, row)
	}
	return result
}

func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		books := GetBooksMap(db)
		fp := path.Join("templates", "index.html")

		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(rw, books); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
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
