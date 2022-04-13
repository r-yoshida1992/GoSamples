package main

import "fmt"

func main() {
	// mapの作成
	m1 := map[string]string{
		"red":  "赤",
		"blue": "青",
	}
	fmt.Println(m1)         // -> map[blue:青 red:赤]
	fmt.Println(m1["red"])  // -> 赤
	fmt.Println(m1["blue"]) // -> 青

	// makeによるマップの生成
	m2 := make(map[string]string)
	m2["apple"] = "リンゴ"
	m2["orange"] = "オレンジ"
	fmt.Println(m2)           // -> map[apple:リンゴ orange:オレンジ]
	fmt.Println(m2["apple"])  // -> リンゴ
	fmt.Println(m2["orange"]) // -> オレンジ

	// 要素の削除
	delete(m2, "apple")
	fmt.Println(m2) // -> map[orange:オレンジ]

	// キーの存在確認
	_, ok := m2["orange"]
	fmt.Println(ok) // -> true
	_, ok2 := m2["apple"]
	fmt.Println(ok2) // -> false

}
