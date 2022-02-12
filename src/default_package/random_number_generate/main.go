package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := generate()
	fmt.Println(num)
}

// 0から9までの乱数を生成
func generate() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}
