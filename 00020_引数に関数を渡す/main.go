package main

import "fmt"

func main() {
	fmt.Println(calc(funcA, 10, 5)) // -> 15
	fmt.Println(calc(funcB, 10, 5)) // -> 5
}

// 関数を引数で受け取る
func calc(fn func(int, int) int, a int, b int) int {
	return fn(a, b)
}

// 足し算
func funcA(a, b int) int {
	return a + b
}

// 引き算
func funcB(a, b int) int {
	return a - b
}
