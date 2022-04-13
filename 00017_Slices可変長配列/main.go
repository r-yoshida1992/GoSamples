package main

import (
	"fmt"
)

func main() {
	// 通常の配列(Arrays)
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a) // -> [1 2 3 4 5 6]

	// Slices
	var s []int
	s = a[0:3]     // aの要素0から3つめの要素までを抜き出す
	fmt.Println(s) // -> [1 2 3]

	s[0] = 99 // スライスは配列の参照であるため、a[0]も上書きされる

	fmt.Println(a) // -> [99 2 3 4 5 6]
	fmt.Println(s) // -> [99 2 3]

	// スライスは以下のようにも取れる
	x := a[:4]     // a[0](最初)から4つめの要素まで
	y := a[2:]     // a[2]から最後まで
	z := a[:]      // a[0](最初)から最後まで
	fmt.Println(x) // -> [99 2 3 4]
	fmt.Println(y) // -> [3 4 5 6]
	fmt.Println(z) // -> [99 2 3 4 5 6]

	// 配列の長さ(len)と容量(cap)
	fmt.Println(len(a)) // -> 6
	fmt.Println(cap(a)) // -> 6

	// スライスの長さ(len)と容量(cap)
	s2 := a[3:4]
	fmt.Println(s2)      // -> [4]
	fmt.Println(len(s2)) // -> 1
	fmt.Println(cap(s2)) // -> 3

	// 組み込みのmakeでのスライス作成
	// make([スライスの型], [長さ(length), [容量(capacity) ※省略可])
	s3 := make([]int, 4, 4)
	fmt.Println(s3) // -> [0 0 0 0]

	// スライスへの要素の追加
	s3 = append(s3, 5)
	fmt.Println(s3) // -> [0 0 0 0 5]
}
