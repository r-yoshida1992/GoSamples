package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// jsonの読み込み
func main() {
	// test.jsonを読み込む
	raw, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	// jsonを配列に格納する
	var persons []Person
	json.Unmarshal(raw, &persons)

	// 繰り返し出力
	for index, person := range persons {
		fmt.Println(strconv.Itoa(index) + " : " + person.Name)
		fmt.Println(strconv.Itoa(index) + " : " + strconv.Itoa(person.Age))
		fmt.Println(strconv.Itoa(index) + " : " + person.Gender)
	}

}

// Person jsonの構造体
type Person struct {
	Name   string
	Age    int
	Gender string
}
