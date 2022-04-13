package main

import (
	"fmt"
	"strconv"
)

func main() {
	person1 := Person{name: "taro", age: 20}
	person2 := Person{"hana", 25}
	person3 := Person{}
	personPointer := &person2
	fmt.Println(person1.name + "は" + strconv.Itoa(person1.age) + "歳です。")             // -> taroは20歳です。
	fmt.Println(person2.name + "は" + strconv.Itoa(person2.age) + "歳です。")             // -> hanaは25歳です。
	fmt.Println(person3.name + "は" + strconv.Itoa(person3.age) + "歳です。")             // -> は0歳です。(空文字と0で初期化)
	fmt.Println(personPointer.name + "は" + strconv.Itoa(personPointer.age) + "歳です。") // -> hanaは25歳です。
	fmt.Println(person1)                                                             // -> {taro 20}
	fmt.Println(person2)                                                             // -> {hana 25}
	fmt.Println(person3)                                                             // -> { 0}
	fmt.Println(personPointer)                                                       // -> &{hana 25}
}

// Person 構造体
type Person struct {
	name string
	age  int
}
