package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	pattern1()
	pattern2()
}

// 現在時刻を使うパターン
func pattern1() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10) // 0から9までの乱数を生成
	fmt.Println(i)
}

// 乱数生成器を使うパターン
func pattern2() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	i := rand.Intn(10) // 0から9までの乱数を生成
	fmt.Println(i)
}
