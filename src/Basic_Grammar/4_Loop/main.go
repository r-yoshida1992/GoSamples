package main

import "fmt"

func main() {
	// forループ
	for i := 0; i < 5; i++ {
		fmt.Printf("forループ%d回目\n", i+1)
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("forループ%d回目\n", i+1)
		if i > 3 {
			fmt.Println("continueでスキップ")
			continue
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("forループ%d回目\n", i+1)
		if i > 3 {
			fmt.Println("breakで抜ける")
			break
		}
	}

	i := 0
	for i < 10 {
		fmt.Println("while的な")
		i++
	}

	// 無限ループ
	for {
		break
	}

	for i, v := range []string{"foo", "bar", "baz"} {
		fmt.Println(i, v)
		// 0 foo
		// 1 bar
		// 2 baz
	}

	for k, v := range map[string]int{"key-1": 100, "key-2": 200, "key-3": 300} {
		fmt.Println(k, v)
		// key-2 200
		// key-3 300
		// key-1 100
	}

}
