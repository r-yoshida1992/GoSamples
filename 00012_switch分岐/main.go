package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	pattern1(rand.Intn(3))
	pattern2(rand.Intn(3))
	pattern3(rand.Intn(3))
}

// 通常のswitch
// ※Goではbreakは書かなくてよい
func pattern1(i int) {
	switch i {
	case 1:
		fmt.Println("1です")
	case 2:
		fmt.Println("2です")
	default:
		fmt.Println("2より大きいです")
	}
}

// 条件の前に変数宣言
func pattern2(i int) {
	switch j := i * 2; j {
	case 2:
		fmt.Println("iの2倍は2です")
	case 4:
		fmt.Println("iの2倍は4です")
	default:
		fmt.Println("iの2倍は4より大きいです")
	}
}

// switchの後を省略するとswitch trueと同義
// if-elseの代わりに使える
func pattern3(i int) {
	switch {
	case i == 1:
		fmt.Println("iは1です")
	case i == 2:
		fmt.Println("iは2です")
	default:
		fmt.Println("iは2より大きいです")
	}
}
