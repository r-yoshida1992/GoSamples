package main

import (
	"strconv"
)

func main() {
	var num int = 1

	// 数値 -> 文字
	var str string = strconv.Itoa(num)

	// 文字 -> 数値
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
}
