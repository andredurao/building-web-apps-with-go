package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		SayHelloName(rw, r)
		return
	}
	http.NotFound(rw, r)
	return
}

func SayHelloName(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
