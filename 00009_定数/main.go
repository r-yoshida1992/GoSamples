package main

import "fmt"

// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使える
func main() {
	const a = 'a'      // character
	const b = "string" // string
	const c = true     // boolean
	const d = 1        // numeric
	fmt.Println(a, b, c, d)
}
