package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

// Action defines a standar funcion signature for us to use wen creating
// controller actions. A controller action is basically just a method attached to
// a controller
type Action func(rw http.ResponseWriter, r *http.Request) error

// This is our Base Controller
type AppController struct{}

type Message struct {
	Title string
	Body  string
}

// Renders the error template on a string to customize the error 500
func RenderErrorPage(err string) string {
	message := Message{"Error", err}
	fp := path.Join("templates", "error.html")
	tmpl, _ := template.ParseFiles(fp)
	buf := new(bytes.Buffer)

	if executeError := tmpl.Execute(buf, message); executeError != nil {
		panic(executeError)
	}
	return buf.String()
}

// The action function helps with error handling in a controller
func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			fmt.Println("--->")
			fmt.Println(err)
			http.Error(rw, RenderErrorPage(err.Error()), 500)
			// http.Error(rw, RenderErrorPage(ex_err), 500)
		}
	})
}
