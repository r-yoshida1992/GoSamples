package main

import (
	"flag"
	"fmt"
)

// 第一引数をコンソールに出力する
func main() {
	flag.Parse()
	_, err := fmt.Println(flag.Args()[0])
	if err != nil {
		panic(err)
	}
}
