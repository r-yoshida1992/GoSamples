package main

import "github.com/gin-gonic/gin"

func main() {
	// GinのEngineインスタンスを取得
	router := gin.Default()

	// パターン指定でHtmlファイルをロードする
	router.LoadHTMLGlob("*.html")

	// Getのハンドラを登録
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"msg": "Hello World",
		})
	})

	// サービス開始
	router.Run(":8080")
}
