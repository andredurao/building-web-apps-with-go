package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: [k] value: [v]", k, strings.Join(v, ""))
	}
	fmt.Fprintf(rw, "Hello world") // Send data to client side
}

func main() {
	http.HandleFunc("/", sayHelloName)       // set router
	err := http.ListenAndServe(":9090", nil) // Set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
