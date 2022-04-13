package main

import "fmt"

func main() {
	var i interface{}

	i = 42
	describe(i)

	i = "hello"
	describe(i)

}

// interface{}は全ての型を受け取れる。Go 1.18ではanyと書ける
func describe(i interface{}) {
	fmt.Printf("type : %T , value : %v\n", i, i)
}
