package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//　0から2までの乱数を生成
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(3)

	if num == 1 {
		fmt.Println("if")
	} else if num == 2 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

}
