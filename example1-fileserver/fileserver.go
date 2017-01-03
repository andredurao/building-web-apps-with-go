package main

// Added the log interface to check what is being served
import (
	"log"
	"net/http"
)

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.ListenAndServe(":8080", loggingHandler(http.FileServer(http.Dir("."))))
}
