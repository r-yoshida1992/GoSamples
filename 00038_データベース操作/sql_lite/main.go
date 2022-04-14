package main

import (
	"database_oparation/base"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// DBに接続
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	// 最後にDBをクローズ
	defer db.Close()

	// DB操作部分については以下の関数内を参照
	base.Operation(db)
}
