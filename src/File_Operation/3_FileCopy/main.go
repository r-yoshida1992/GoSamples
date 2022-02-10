package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 引数チェック
	if len(os.Args) < 3 {
		fmt.Println("引数1と引数2を指定してください。")
	}

	// コピー元とコピー先のファイル名
	sourceFileName := os.Args[1]
	copyFileName := os.Args[2]

	src, err := os.Open(sourceFileName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(copyFileName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	// コピー処理
	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}

	fmt.Printf("「%s」を「%s」にコピーしました。", sourceFileName, copyFileName)
}
