package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// jsonの読み込み
func main() {
	// test.jsonを読み込む
	raw, err := ioutil.ReadFile("./test.json")
	if err != nil {
		panic(err)
	}

	// jsonを配列に格納する
	var persons []Person
	json.Unmarshal(raw, &persons)

	// jsonの中身を出力
	for _, person := range persons {
		fmt.Println(person)
	}

}

// Person jsonの構造体
// 要素は必ず大文字で始める
type Person struct {
	Name   string
	Age    int
	Gender string
}
