package main

import (
	"errors"
	"github.com/unrolled/render"
	"html/template"
	"net/http"
	"path"
)

type MyController struct {
	AppController
	*render.Render
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func (c *MyController) Status(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Status": "RUNNING"})
	return nil
}

func (c *MyController) FaultAction(rw http.ResponseWriter, r *http.Request) error {
	return errors.New("some fault in action")
}

func (c *MyController) Home(rw http.ResponseWriter, r *http.Request) error {
	message := Message{"Hello", "This is the home action"}
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return nil
	}

	if err := tmpl.Execute(rw, message); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return nil
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}
	http.Handle("/", c.Action(c.Index))
	http.Handle("/status", c.Action(c.Status))
	http.Handle("/home", c.Action(c.Home))
	http.Handle("/fault", c.Action(c.FaultAction))
	http.ListenAndServe(":8080", nil)
}
