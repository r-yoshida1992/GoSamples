package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(writer, "Hello World")
	if err != nil {
		return
	}
}
