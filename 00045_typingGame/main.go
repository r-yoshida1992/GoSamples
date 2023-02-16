package main

import "fmt"

var marks = []string{"|", "/", "-", "\\"}

func mark(i int) string {
	return marks[i%4]
}

func main() {
	demo1()
}

func demo1() {
	cnt := 5000000000
	for i := 1; i <= cnt; i++ {
		if i%(cnt/100) == 0 {
			p := i / (cnt / 100)
			fmt.Printf("\rLoading: %s %2d%% (%d)", mark(p), p, i)
		}
	}
	fmt.Printf("\r                                                    ")
	fmt.Printf("\rDone.")
}

func demo2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\raaaa")
	}
}
