package main

import "fmt"

func main() {
	values := []string{"a", "b", "c"}
	for i, v := range values {
		fmt.Println(string(i) + v)
	}
}
