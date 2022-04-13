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

	// 条件の前に変数宣言もできる
	if num2 := rand.Intn(4); num2 < 2 {
		fmt.Println("num2は2より小さい")
	} else {
		fmt.Println("num2は2以上")
	}

}
