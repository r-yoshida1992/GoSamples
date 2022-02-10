package main

import "fmt"

// 変数定義パターン
func main() {
	// 変数を宣言してから代入
	var v1 string
	v1 = "v1"

	// 変数の宣言と代入を同時に行う(型指定)
	var v2 string = "v2"

	// 型を省略
	var v3 = "v3"

	// varを省略
	v4 := "v4"

	// 複数宣言(型同一)
	var v5, v6 string
	v5, v6 = "v5", "v6"

	// var省略(型同一)
	var v7, v8 string = "v7", "v8"

	// 複数宣言(型違い)
	var (
		v9  int // 符号付き整数
		v10 int8
		v11 int16
		v12 int32
		v13 int64
		v14 rune // int32のエイリアス
		v15 uint // 符号なし整数
		v16 uint8
		v17 uint16
		v18 uint32
		v19 uint64
		v20 uintptr // ポインタの値
		v21 byte    // uint8のエイリアス
		v22 float32 // 浮動小数点
		v23 float64
		v24 bool      // 真偽値
		v25 complex64 // 複素数
		v26 complex128
		v27 string // 文字列
	)
	v9 = 9
	v10 = 10
	v11 = 11
	v12 = 12
	v13 = 13
	v14 = 14
	v15 = 15
	v16 = 16
	v17 = 17
	v18 = 18
	v19 = 19
	v20 = 20
	v21 = 21
	v22 = 22
	v23 = 23
	v24 = true
	v25 = 25 + 1i
	v26 = 26 + 1i
	v27 = "v27"

	// 複数宣言(型省略)
	var (
		v28 = 28
		v29 = "v29"
	)

	// 複数宣言(型省略)
	var v30, v31 = "v30", 31

	// 出力 -> v1 v2 v3 v4 v5 v6 v7 v8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 true (25+1i) (26+1i) v27 28 v29 v30 31
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31)
}
