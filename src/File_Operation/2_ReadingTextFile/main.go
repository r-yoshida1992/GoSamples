package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// sample.txtの中身を読み込む
	filename := "./sample.txt"
	fileText, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした")
		return
	}
	fmt.Println(string(fileText))
}
