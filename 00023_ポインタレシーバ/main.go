// ポインタレシーバを使うことでレシーバが指す先の変数を変更できる
package main

import "fmt"

type Person struct {
	name string
	age  int
}

// Increment1 ポインタレシーバ未使用(値レシーバ)。p.ageを更新できない
func (p Person) Increment1() {
	p.age += 1
}

// Increment2 ポインタレシーバを使用。p.ageを更新できる
func (p *Person) Increment2() {
	p.age += 1
}

func main() {
	p := Person{"taro", 20}
	p.Increment1()     // 更新されない
	fmt.Println(p.age) // -> 20
	p.Increment2()     // 更新される
	fmt.Println(p.age) // -> 21
}
