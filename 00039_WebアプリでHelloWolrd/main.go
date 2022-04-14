package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil) // ポート8080で受け取る
}

func handler(writer http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(writer, "Hello World") // HelloWorldをwriterに書き込む
	if err != nil {
		return
	}
}
