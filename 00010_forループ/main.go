package main

import "fmt"

func main() {
	// 通常のfor
	for i := 0; i < 5; i++ {
		fmt.Printf("forループ%d回目\n", i+1)
	}

	// continueでスキップ
	for i := 0; i < 5; i++ {
		fmt.Printf("forループ%d回目\n", i+1)
		if i > 3 {
			fmt.Println("continueでスキップ")
			continue
		}
	}

	// breakで抜ける
	for i := 0; i < 5; i++ {
		fmt.Printf("forループ%d回目\n", i+1)
		if i > 3 {
			fmt.Println("breakで抜ける")
			break
		}
	}

	// while的なやつ
	i := 0
	for i < 10 {
		fmt.Println("while的な")
		i++
	}

	// 無限ループ
	for {
		break
	}

	// foreach的なやつ
	for i, v := range []string{"foo", "bar", "baz"} {
		fmt.Println(i, v)
		// 0 foo
		// 1 bar
		// 2 baz
	}

	// mapでも使える
	for k, v := range map[string]int{"key-1": 100, "key-2": 200, "key-3": 300} {
		fmt.Println(k, v)
		// key-2 200
		// key-3 300
		// key-1 100
	}

}
