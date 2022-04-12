package main

import "fmt"

func main() {
	result1 := pattern1(3, 4)
	fmt.Println(result1) // -> 7
	result2, result3 := pattern2(3, 4)
	fmt.Println(result2) // -> 4
	fmt.Println(result3) // -> 3
	result4, result5 := pattern3(7, 8)
	fmt.Println(result4) // -> 7
	fmt.Println(result5) // -> 8
}

// 戻り値が一つのパターン
func pattern1(a int, b int) int {
	return a + b
}

// 戻り値が複数のパターン (※二つ以上でも可)
func pattern2(a int, b int) (int, int) {
	return b, a
}

// 引数に名前がつくパターン (※returnに記述しなくてもよい)
func pattern3(a int, b int) (x, y int) {
	x = a
	y = b
	return
}
