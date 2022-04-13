package main

import "fmt"

func main() {
	// 配列の宣言
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // -> Hello World
	fmt.Println(a)          // -> [Hello World]

	b := [3]int{1, 2, 3}
	fmt.Println(b) // -> [1 2 3]
}
