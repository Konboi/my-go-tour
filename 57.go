package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	var h Hello
	fmt.Println(h)
	http.ListenAndServe("localhost:4000", h)
}
