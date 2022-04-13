package main

import "fmt"

// InterfaceSample インターフェース
type InterfaceSample interface {
	Method()
}

type T struct {
	S string
}

// Method インターフェースの実装
func (t T) Method() {
	fmt.Println(t.S, "T!!!")
}

type U struct {
	S string
}

// Method インターフェースの実装
func (u U) Method() {
	fmt.Println(u.S, "U!!!")
}

func main() {
	var a1 InterfaceSample = T{"hello"}
	a1.Method() // -> hello T!!!
	var a2 InterfaceSample = U{"hello"}
	a2.Method() // -> hello U!!!
}
