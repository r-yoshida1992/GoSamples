package main

import (
	"fmt"
	"time"
)

func main() {
	// そのまま表示
	fmt.Println(time.Now())
	// フォーマット
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
}
