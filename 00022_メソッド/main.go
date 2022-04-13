package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Vertexのメソッドとして使える
func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y) // 平方根を返す
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.abs()) // -> (3*3) + (4*4)の平方根「5」を出力
}
