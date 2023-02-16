package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
)

func main() {
	url := "https://www.smule.com/Ren_ren_z/performances/json?offset=0"
	scrapeText(url)
}

func scrapeText(url string) {

	// Getリクエストでレスポンス取得
	res, _ := http.Get(url)
	defer res.Body.Close()

	// Body内を読み取り
	buffer, _ := ioutil.ReadAll(res.Body)

	// 文字コードを判定
	detector := chardet.NewTextDetector()
	detectResult, _ := detector.DetectBest(buffer)
	fmt.Println(detectResult.Charset)
	// => UTF-8

	// 文字コードの変換
	bufferReader := bytes.NewReader(buffer)
	reader, _ := charset.NewReaderLabel(detectResult.Charset, bufferReader)

	// HTMLをパース
	document, _ := goquery.NewDocumentFromReader(reader)

	result := document.Find("body").Text()
	fmt.Println(result)

}
