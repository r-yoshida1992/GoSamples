package main

import "fmt"

func main() {
	pattern1()
	pattern2()
}

// deferをつけると関数の終わりに実行される
func pattern1() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// deferを複数回呼ぶと最後の呼び出しから実行される
func pattern2() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		// 9から出力され、最後に0が出力される
	}

	fmt.Println("done")
}
