package main

import "fmt"

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),  // iの数だけ加算する
			neg(-i), // iの数だけ減算する
		)
	}
}

// クロージャーを返す
func adder() func(int) int {
	sum := 0 // line9,10で渡された引数はこのsum変数へバインドされる
	return func(x int) int {
		sum += x
		return sum
	}
}
