package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 同じディレクトリに「test.txt」を作成する。
	filename := "./test.txt"
	// ファイル名、テキストの中身、権限の順で指定
	err := ioutil.WriteFile(filename, []byte("test"), 0777)
	// エラー処理
	if err != nil {
		fmt.Printf("%sへの書き込みに失敗しました。", filename)
		return
	}
	fmt.Printf("%sへの書き込みに成功しました。", filename)
}
