package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) // iがstringであるため、helloが代入される
	fmt.Println(s)  // -> hello

	s, ok := i.(string) // iがstringであるため、okにはtrueが代入される
	fmt.Println(s, ok)  // -> hello true

	f, ok := i.(float64) // iがstringであるため、okにはfalseが代入される
	fmt.Println(f, ok)   // -> 0 false

	f = i.(float64) // iがstringであるため、panicが起こる
	fmt.Println(f)
}
