package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World Go!")
	})
	http.ListenAndServe(":1337", nil)
}
