package main

import (
	"e_learning_sample/app/server"
	"e_learning_sample/config"
	"flag"
)

func main() {
	// 引数を取得
	flag.Parse()

	// 引数から環境を設定
	config.Init(flag.Args()[0])
	conf := config.GetConfig()

	// サービス開始
	r := server.GetRouter()
	r.Run(":" + conf.GetString("port"))
}
