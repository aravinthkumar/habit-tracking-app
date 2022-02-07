package main

import (
	"fmt"
	"net/http"
)

type habit string
type tracker int

func (h habit) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello here")
}

func (t tracker) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, 1)
}

func main() {
	var h habit
	var t tracker
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	})
	http.Handle("/h", h)
	http.Handle("/t", t)
	http.ListenAndServe(":8080", nil)
}
