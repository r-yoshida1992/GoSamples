package main

import "fmt"

func main() {
	do(1)
	do("hello")
	do(true)
}

func do(i any) {
	switch i.(type) {
	case int:
		fmt.Println("数値です。")
	case string:
		fmt.Println("文字列です。")
	default:
		fmt.Println("数値でも文字列でもないです。なんこれ。")
	}
}
