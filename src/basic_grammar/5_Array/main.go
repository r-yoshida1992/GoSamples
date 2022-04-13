package main

import "fmt"

func main() {

	// ノーマルな配列の宣言
	values := []string{"a", "b", "c"}
	for i, v := range values {
		fmt.Println(i, v)
	}

	// 配列の追加
	values = append(values, "d")
	for i, v := range values {
		fmt.Println(i, v)
	}
}
