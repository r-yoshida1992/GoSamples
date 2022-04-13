package main

import "fmt"

func main() {
	a := 1
	b := &a
	c := *b
	fmt.Println(a) // -> 1
	fmt.Println(b) // -> aのメモリアドレス
	fmt.Println(c) // -> 1 (bの指す先の変数)
}
